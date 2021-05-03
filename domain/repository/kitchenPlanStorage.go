package repository

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type KitchenPlanStorage interface {
	ingredientReceiver
	receiptReceiver
}

type receiptReceiver interface {
	GetReceipts() ([]entity.Receipt, error)
	GetReceipt(id int) (*entity.Receipt, error)
}

type ingredientReceiver interface {
	GetIngredients(title string) ([]entity.Ingredient, error)
	GetIngredient(id int) (*entity.Ingredient, error)
}
