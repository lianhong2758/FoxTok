package sq

import (
	"FoxTok/config"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config.DB.User, config.Config.DB.Pass, config.Config.DB.Name, config.Config.DB.Port)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return
	}
	db.AutoMigrate(new(User), new(PassWord), new(Video), new(IsFavorite), new(Comment), new(Message), new(Follower))
	logrus.Info("已连接PQ DB,初始化完成")
	db.Where("", 1).Delete("", nil)
}
