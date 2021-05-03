package repository

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type ReceiptReceiver interface {
	GetReceipts(selection entity.Selection) ([]entity.Receipt, error)
	GetReceipt(id int) (*entity.Receipt, error)
}
