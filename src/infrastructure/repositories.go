package infrastructure

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/src/config"
	"github.com/softree-group/kitchen-plan-backend/src/domain/repository"
	"github.com/softree-group/kitchen-plan-backend/src/infrastructure/persistence"
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

	receiptReceiver := persistence.NewReceiptReceiver(conn)
	receiptReceiver.Prepare()

	ingredientReceiver := persistence.NewIngredientsReceiver(conn)
	ingredientReceiver.Prepare()

	return &repository.Repositories{
		IngredientReceiver: ingredientReceiver,
		ReceiptReceiver:    receiptReceiver,
	}
}
