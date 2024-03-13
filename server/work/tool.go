package work

import (
	"FoxTok/server/template"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

type RequestType interface {
	template.UserLoginRequest | template.UserRegisterRequest|template.FeedRequest
}

func UnmarshalRequest[T RequestType, PT interface{*T ;proto.Message}](ctx *gin.Context) (t *T, err error) {
	var req  PT = new(T)
	body, _ := ctx.GetRawData()
	err = proto.Unmarshal(body, req)
	if err != nil {
		ctx.Abort()
	}
	return req,nil
}
