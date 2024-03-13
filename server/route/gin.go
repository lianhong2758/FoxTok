package route

import (
	"FoxTok/config"
	"FoxTok/server/work"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()                                    //初始化
	r.GET("/file", work.GetFile)                          //获取文件等
	dy := r.Group("/douyin")                              //鉴权无
	dy.GET("/user", work.UserInfo)            //获取用户信息
	dy.POST("/user/register/", work.Register)             // 用户注册接口
	dy.POST("/user/login/", work.Login)                   //用户登录接口
	dy.GET("/feed", work.Feed)                            //视频流
	dy.POST("/publish/action", func(ctx *gin.Context) {}) //投稿
	dy.GET("publish/list", func(ctx *gin.Context) {})     //发布列表
	r.Run(fmt.Sprintf("0.0.0.0:%d", config.Config.Gin.Port))
}
