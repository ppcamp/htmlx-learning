package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func HtmlxTodos(c *gin.Context) {
	c.HTML(200, "routes/htmlx/todos.component.tmpl", "")
}

func HtmlxVideos(c *gin.Context) {
	time.Sleep(2 * time.Second)
	c.HTML(200, "routes/htmlx/videos.component.tmpl", "")
}
