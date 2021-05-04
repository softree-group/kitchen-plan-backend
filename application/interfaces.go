package application

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type IngredientReceiver interface {
	Filter(title string) ([]entity.Ingredient, error)
	Receive(id int) (*entity.Ingredient, error)
}

type ReceiptReceiver interface {
	Filter(filter entity.ReceiptFilter) ([]entity.Receipt, error)
	Receive(id int) (*entity.Receipt, error)
}
