package infrastructure

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/config"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
	"github.com/softree-group/kitchen-plan-backend/infrastructure/persistence"
	"github.com/spf13/viper"
)

func New() *repository.Repositories {
	persistence.Migrate()

	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     viper.GetString(config.DBHost),
			Port:     uint16(viper.GetInt(config.DBPort)),
			Database: viper.GetString(config.DBName),
			User:     viper.GetString(config.DBUser),
			Password: viper.GetString(config.DBPassword),
		},
		MaxConnections: 100,
	})

	if err != nil {
		logrus.Fatal("Fail to create db repository: ", err)
	}

	return &repository.Repositories{
		IngredientReceiver: persistence.NewIngredientsReceiver(conn),
		ReceiptReceiver:    persistence.NewReceiptReceiver(conn),
	}
}
