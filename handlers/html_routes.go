package handlers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.HTML(200, "routes/index.tmpl", "")
}

func Watch(c *gin.Context) {
	c.HTML(200, "routes/watch/index.tmpl", gin.H{"id": c.Param("video")})
}

func NotFound(c *gin.Context) {
	c.HTML(200, "routes/not-found.tmpl", "")
}
