package sq

import (
	"FoxTok/config"
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%ssslmode=disable", config.Config.DB.Host, config.Config.DB.Port, config.Config.DB.User, config.Config.DB.Pass, config.Config.DB.Name)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = db.Ping()
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("已连接PQ DB")
}