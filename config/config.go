package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// 单例
var Config *config

type config struct {
	DB   *DB   `yaml:"db"`
	Gin *Gin `yaml:"gin"`
}
type DB struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}
type Gin struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	VipVideo string `yaml:"vip_video"`
	Video    string `yaml:"video"`
}

func ReadConfig(path string) {
	Config = new(config)
	data, err := os.ReadFile(path)
	if err != nil {
		Config.DB = &DB{
			Host: "localhost",
			Port: 5432,
			User: "root",
			Pass: "2758",
			Name: "go",
		}
		Config.Gin = & Gin{
			Host:     "localhost",
			Port:      9600,
			VipVideo: "http://localhost:8080/vip_video",
			Video:    "http://localhost:8080/ideo",
		}
		data, _ = yaml.Marshal(Config)
		err = os.WriteFile(path, data, 0644)
		if err != nil {
			logrus.Fatalln(err)
		}
		logrus.Info("填写config.yml后再次运行")
		os.Exit(0)
	}
	err = yaml.Unmarshal(data, Config)
	if err != nil {
		logrus.Fatalln(err)
	}
	return
}
