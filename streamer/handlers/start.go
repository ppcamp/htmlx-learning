package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/movies-to-watch/streamer/config"
	"github.com/ppcamp/movies-to-watch/streamer/ffmpeg"
)

func Stream(c *gin.Context) {
	playlist := c.Param("name")

	if !fileExists("", playlist, config.PlaylistExt()) {
		c.JSON(404, gin.H{"status": "404", "message": "Not Found. Request to watch it first."})
		return
	}

	c.JSON(200, gin.H{
		"status": "200",
		"file":   fmt.Sprintf("/playlists/%s", ffmpeg.DefaultFileName(playlist)),
	})
}

func Start(c *gin.Context) {
	playlist := c.Param("name")

	if !fileExists(config.VideosFolder(), playlist, config.VideoFileExt()) {
		c.JSON(404, gin.H{"status": "404", "message": "Not Found"})
		return
	}

	if fileExists(config.PlaylistExt(), playlist, config.PlaylistExt()) {
		c.JSON(200, gin.H{"status": "200", "message": "OK"})
		return
	}

	q := ffmpeg.DefaultQuery(playlist)
	exec := ffmpeg.NewExecutor(q)
	if err := exec.Run(); err != nil {
		c.JSON(500, gin.H{"status": "500", "message": "Fail to init the executor"})
		return
	}

	c.JSON(200, gin.H{"status": "200", "message": "Converting"})
}
