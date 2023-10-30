package handlers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/movies-to-watch/common/files"
	log "github.com/sirupsen/logrus"
)

func Stream(c *gin.Context) {
	playlist := c.Param("name")
	if !fileExists(playlist) {
		// should initialize the executor
		c.JSON(404, gin.H{"status": "404", "message": "Not Found. Initialize the executor"})
		return
	}

	// available to stream
	c.JSON(200, gin.H{"status": "200", "file": fmt.Sprintf("/playlists/%s-playlist.m3u8", playlist)})
}

func Start(c *gin.Context) {
	playlist := c.Param("name")
	if !fileExists(playlist) {
		// should initialize the executor
		c.JSON(404, gin.H{"status": "404", "message": "Not Found"})
		return
	}

	// available to stream
	c.JSON(200, gin.H{"status": "200", "message": "OK"})
}

func fileExists(n string) bool {
	playlists, err := files.GlobSuffix(".m3u8")
	if err != nil {
		log.WithError(err).Warn("failed to get playlists")
		return false
	}

	for _, v := range playlists {
		if strings.Contains(v, n) {
			return true
		}
	}
	return false
}
