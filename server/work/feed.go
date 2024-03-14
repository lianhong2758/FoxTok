package work

import (
	"FoxTok/server/sq"
	"FoxTok/server/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func Feed(ctx *gin.Context) {
	ltime, token := ctx.Query("latest_time"), ctx.Query("token")
	result := new(template.FeedResponse)
	result.StatusCode = 0

	switch token {
	case "":
		//游客
		result.StatusMsg = proto.String("Sucess")
		videos := []*template.Video{}
		sltime, _ := strconv.ParseInt(ltime, 10, 64)
		v, t := sq.GetNextVideo(sltime, 0)
		videos = append(videos, v)
		result.VideoList = videos
		result.NextTime = proto.Int64(t)
	default:
	}
	ctx.JSON(200, result)
}
