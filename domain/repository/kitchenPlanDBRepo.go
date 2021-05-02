package repository

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type KitchenPlanDBRepo interface {
	getIngredients
	getIngredient
	getReceipts
	getReceipt
}

type getIngredients interface {
	GetIngredients(title string) ([]entity.Ingredient, error)
}

type getIngredient interface {
	GetIngredient(id int) (*entity.Ingredient, error)
}

type getReceipts interface {
	GetReceipts() ([]entity.Receipt, error)
}

type getReceipt interface {
	GetReceipt(id int) (*entity.Receipt, error)
}
