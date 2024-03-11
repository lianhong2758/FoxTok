package route

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()                                    //初始化
	dy := r.Group("/douyin", func(ctx *gin.Context) {})   //鉴权
	dy.GET("/user", func(ctx *gin.Context) {})            //获取用户信息
	dy.POST("/user/register", func(ctx *gin.Context) {})  // 用户注册接口
	dy.POST("/user/login", func(ctx *gin.Context) {})     //用户登录接口
	dy.GET("/feed", func(ctx *gin.Context) {})            //视频流
	dy.POST("/publish/action", func(ctx *gin.Context) {}) //投稿
	dy.GET("publish/list", func(ctx *gin.Context) {})     //发布列表
	r.Run("0.0.0.0:9600")
}
