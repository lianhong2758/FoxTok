package sq

import (
	"FoxTok/config"
	"fmt"
)

const Http = "http://"

func GetFileUri() string {
	return fmt.Sprintf("%s%s:%d/file?path=", Http, config.Config.Gin.Host, config.Config.Gin.Port)
}

func GetVipFilePath() string {
	return config.Config.Gin.VipVideo
}
func GetFilePath() string {
	return config.Config.Gin.Video
}
