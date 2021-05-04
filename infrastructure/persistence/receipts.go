package persistence

import (
	"github.com/jackc/pgx"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
)

type ReceiptReceiver struct {
	db *pgx.ConnPool
}

func (r ReceiptReceiver) GetReceipts(selection *entity.ReceiptFilter) ([]entity.Receipt, error) {
	panic("implement me")
}

func (r ReceiptReceiver) GetReceipt(id int) (*entity.Receipt, error) {
	panic("implement me")
}

var _ repository.ReceiptReceiver = &ReceiptReceiver{}

func NewReceiptReceiver(db *pgx.ConnPool) *ReceiptReceiver {
	return &ReceiptReceiver{db}
}