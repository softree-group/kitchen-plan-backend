package config

import "github.com/spf13/viper"

func defaults() {
	mainDefaults()
	logDefaults()
	dbDefaults()
	staticDefaults()
}

func mainDefaults() {
	viper.SetDefault(IP, "127.0.0.1")
	viper.SetDefault(Port, "8000")
}

func logDefaults() {
	viper.SetDefault(LogFile, "log.log")
	viper.SetDefault(LogLevel, "debug")
}

func dbDefaults() {
	viper.SetDefault(DBName, "")
	viper.SetDefault(DBHost, "127.0.0.1")
	viper.SetDefault(DBPort, "5432")
	viper.SetDefault(DBUser, "")
	viper.SetDefault(DBPassword, "")
	viper.SetDefault(DBMigrations, "migrations")
}

func staticDefaults() {
	viper.SetDefault(StaticStorageRoot, "")
}
