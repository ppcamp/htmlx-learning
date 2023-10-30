package handlers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	data := MapBase(c)
	c.HTML(200, "routes/index.tmpl", data)
}

func Watch(c *gin.Context) {
	data := MapBase(c)
	c.HTML(200, "routes/watch/index.tmpl", data)
}

func NotFound(c *gin.Context) {
	c.HTML(200, "routes/not-found.tmpl", "")
}
