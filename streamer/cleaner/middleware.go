package cleaner

import (
	"github.com/gin-gonic/gin"
)

// TODO: add support to recognize fetch from static folder
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		v := c.Param("video")
		if v != "" {
			Push(v)
		}

		c.Next()
	}
}
