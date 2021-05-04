package repository

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type ReceiptReceiver interface {
	Filter(selection *entity.ReceiptFilter) ([]entity.Receipt, error)
	Receive(id int) (*entity.Receipt, error)
}
