package persistence

import (
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
	panic("implement me")
}

func (r ReceiptReceiver) Receive(id int) (*entity.Receipt, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { endTx(tx, err) }()
	var receipt entity.Receipt
	var steps string

	err = tx.QueryRow(sqlReceiveReceipt, id).Scan(&receipt.Id, &receipt.Type, &receipt.Image, &receipt.Title,
		&steps, &receipt.TimeToCook)
	if err != nil {
		return nil, err
	}
	receipt.Steps = strings.Split(steps, "|~~~|")

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
