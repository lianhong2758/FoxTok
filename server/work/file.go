package work

import "github.com/gin-gonic/gin"

func GetFile(ctx *gin.Context) {
	path := ctx.Query("path")
	if path == "" {
		ctx.Abort()
	}
	ctx.File(path)
}
