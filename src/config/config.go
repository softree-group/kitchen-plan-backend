package config

import (
	"flag"
	"fmt"
	"github.com/mark-by/logutils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configPath string

func initFlags() {
	flag.StringVar(&configPath, "s", "config.yaml", "путь к файлу с настройками")

	flag.Parse()
}

func InitConfig() {
	initFlags()

	viper.SetConfigFile(configPath)

	defaults()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			logrus.Fatalf("Невалидный синтаксис")
		} else {
			_ = viper.WriteConfig()
			fmt.Printf("Файл не найден. По пути '%s' записан шаблон\n", "config.yaml")
		}
	}

	viper.AllKeys()
	logutils.InitLogrus(viper.GetString(LogFile), viper.GetString(LogLevel))
}
