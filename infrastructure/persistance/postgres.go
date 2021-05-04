package persistance

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/spf13/viper"
	"time"
)

type PostgresStorage struct {
	conn   *sql.DB
	timing time.Duration
}

func NewPostgresStorage(conn *sql.DB) *PostgresStorage {
	var timing time.Duration = 5
	if viper.GetDuration("sql.timing") != 0 {
		timing = viper.GetDuration("sql.timing")
	}

	return &PostgresStorage{conn: conn, timing: timing * time.Second}
}

func (storage *PostgresStorage) GetReceipts(selection entity.Selection) ([]entity.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), storage.timing)
	defer cancel()

	tx, err := storage.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = tx.Rollback(); err != nil {
			logrus.Info(err)
			return
		}
	}()

	query, args := chooseSql(selection)
	out, err := tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = out.Close(); err != nil {
			logrus.Info(err)
			return
		}
	}()

	receipts := make([]entity.Receipt, 0)
	for out.Next() {
		receipt := entity.Receipt{}

		err = out.Scan(&receipt.Id, &receipt.Image, &receipt.Title, &receipt.TimeToCook, &receipt.Type)
		if err != nil {
			return nil, err
		}

		receipts = append(receipts, receipt)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return receipts, nil
}

const (
	allReceipts      = "select r.id, image, title, time, type from recipes r"
	allLimitReceipts = allReceipts + " limit $1"

	typeReceipts           = allReceipts + " where type=$1"
	typeLimitReceipts      = typeReceipts + " limit $2"
	typeTitleReceipts      = typeReceipts + " and title=$2"
	typeTitleLimitReceipts = typeTitleReceipts + " limit $3"

	titleReceipts      = allReceipts + " where title=$1"
	titleLimitReceipts = titleReceipts + " limit $2"

	ingredientsReceipts          = allReceipts +
		" join recipes_ingridients ri on r.id=ri.recept_id" +
		" where ri.ingridient_id in $1"
	ingredientsLimitReceipts     = ingredientsReceipts + " limit $2"
	ingredientsTypeReceipts      = ingredientsReceipts + " and type=$2"
	ingredientsTypeLimitReceipts = ingredientsTypeReceipts + " limit $3"
)

func (storage *PostgresStorage) GetReceipt(id int) (*entity.Receipt, error) {
	return nil, nil
}

func (storage *PostgresStorage) GetIngredients(title string) ([]entity.Ingredient, error) {
	return nil, nil
}

func (storage *PostgresStorage) GetIngredient(id int) (*entity.Ingredient, error) {
	return nil, nil
}
