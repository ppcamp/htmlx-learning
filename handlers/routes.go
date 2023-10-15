package handlers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.HTML(200, "index.tmpl", "")
}

func Bookmarks(c *gin.Context) {
	c.HTML(200, "bookmarks/index.tmpl", "")
}

func Search(c *gin.Context) {
	c.HTML(200, "index.tmpl", "")
}
