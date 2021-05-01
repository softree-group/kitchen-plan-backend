package config

import "github.com/spf13/viper"

func defaults() {
	mainDefaults()
	logDefaults()
}

func mainDefaults() {
	viper.SetDefault(IP, "127.0.0.1")
	viper.SetDefault(Port, "8989")
}

func logDefaults() {
	viper.SetDefault(LogFile, "log.log")
	viper.SetDefault(LogLevel, "debug")
}
