package work

import (
	"FoxTok/server/template"
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func Register(ctx *gin.Context) {
	name, pass := ctx.Query("username"), ctx.Query("password")
	fmt.Println(name, pass)
	ctx.JSON(200, template.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  proto.String("Sucess"),
		UserId:     123,
		Token:      "sadad",
	})
}
func Login(ctx *gin.Context) {

}

func UserInfo(ctx *gin.Context) {
	id, token := ctx.Query("user_id"), ctx.Query("token")
	fmt.Println(id,token)
	ctx.JSON(200, template.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  proto.String("Sucess"),
		User: &template.User{
			Id:   123,
			Name: "123",
		},
	})
}
