package persistence

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/config"
	"github.com/spf13/viper"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() {
	m, err := migrate.New(
		"file://infrastructure/persistent/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			viper.GetString(config.DBUser),
			viper.GetString(config.DBPassword),
			viper.GetString(config.DBHost),
			viper.GetString(config.DBPort),
			viper.GetString(config.DBName),
		),
	)
	if err != nil {
		logrus.Fatal("Fail to connect to database: ", err)
	}
	defer m.Close()

	err = m.Up()
	switch err {
	case nil:
		logrus.Info("Migrate status: migrations applied")
	case migrate.ErrNoChange:
		logrus.Info("Migrate status: no changes")
	default:
		logrus.Fatal("Fail to apply migrations: ", err)
	}
}

