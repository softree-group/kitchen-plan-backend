package application

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type appReceiptReceiver interface {
	GetAllReceipts() ([]entity.Receipt, error)
	GetReceipt(id int) (*entity.Receipt, error)
}
