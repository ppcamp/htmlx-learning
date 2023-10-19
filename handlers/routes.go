package handlers

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "routes/index.tmpl", "")
}

func Bookmarks(c *gin.Context) {
	c.HTML(200, "routes/bookmarks/index.tmpl", "")
}

func Watch(c *gin.Context) {
	c.HTML(200, "routes/bookmarks/index.tmpl", "")
}

func Search(c *gin.Context) {
	c.HTML(200, "routes/index.tmpl", "")
}

func Todos(c *gin.Context) {
	c.HTML(200, "routes/htmlx/todos.tmpl", "")
}

func HtmlxVideos(c *gin.Context) {
	c.HTML(200, "routes/htmlx/videos.tmpl", "")
}

func Videos(c *gin.Context) {
	c.HTML(200, "routes/videos/index.tmpl", "")
}

func NotFound(c *gin.Context) {
	c.HTML(200, "routes/not-found.tmpl", "")
}
