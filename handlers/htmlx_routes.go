package handlers

import "github.com/gin-gonic/gin"

func Todos(c *gin.Context) {
	c.HTML(200, "routes/htmlx/todos.component.tmpl", "")
}

func HtmlxVideos(c *gin.Context) {
	c.HTML(200, "routes/htmlx/videos.component.tmpl", "")
}
