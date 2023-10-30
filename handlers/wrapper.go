package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ppcamp/htmlx-movies-to-watch/utils/files"
)

func MapBase(c *gin.Context) gin.H {
	param := make(map[string]string)
	for _, v := range c.Params {
		param[v.Key] = v.Value
	}

	videos, err := files.GlobSuffix("/mnt/d/transfer/filmes", ".mp4", ".mkv")
	if err != nil {
		c.AbortWithError(500, err)
		panic(err)
	}

	fmt.Println("Videos: ", videos)

	return gin.H{"Params": param, "Videos": videos}
}
