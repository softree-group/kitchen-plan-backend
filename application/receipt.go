package application

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type AppReceiptReceiver interface {
	GetAllReceipts(selection entity.Selection) ([]entity.Receipt, error)
	GetReceipt(id int) (*entity.Receipt, error)
}
