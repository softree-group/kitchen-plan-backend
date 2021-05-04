package persistence

import (
	"database/sql"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
	"strings"
)

const (
	sqlReceiveReceipt = "receiveReceipt"
)

type ReceiptReceiver struct {
	db *pgx.ConnPool
}

func (r ReceiptReceiver) Filter(filter *entity.ReceiptFilter) ([]entity.Receipt, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { endTx(tx, err) }()

	sqlFilter, sqlArgs := genSQLFilter(filter)

	rows, err := tx.Query(sqlFilter, sqlArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []entity.Receipt

	for rows.Next() {
		receipt := entity.Receipt{}
		timeToCook := sql.NullInt32{}
		err = rows.Scan(&receipt.Id, &receipt.Title, &receipt.Image, &timeToCook)
		if err != nil {
			return nil, err
		}
		if timeToCook.Valid {
			receipt.TimeToCook = int(timeToCook.Int32)
		}
		recipes = append(recipes, receipt)
	}

	return recipes, nil
}

func (r ReceiptReceiver) Receive(id int) (*entity.Receipt, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { endTx(tx, err) }()
	var receipt entity.Receipt
	var steps string

	timeToCook := sql.NullInt32{}

	err = tx.QueryRow(sqlReceiveReceipt, id).Scan(&receipt.Id, &receipt.Type, &receipt.Image, &receipt.Title,
		&steps, &timeToCook)
	if err != nil {
		return nil, err
	}
	receipt.Steps = strings.Split(steps, "|~~~|")

	if timeToCook.Valid {
		receipt.TimeToCook = int(timeToCook.Int32)
	}

	return &receipt, nil
}

func (r ReceiptReceiver) Prepare() {
	if _, err := r.db.Prepare(sqlReceiveReceipt,
		"select id, type, image, title, steps, time_to_cook from recipes where id = $1"); err != nil {
		logrus.Fatal(err)
	}
}

var _ repository.ReceiptReceiver = &ReceiptReceiver{}

func NewReceiptReceiver(db *pgx.ConnPool) *ReceiptReceiver {
	return &ReceiptReceiver{db}
}
