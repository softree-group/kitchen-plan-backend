package application

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type Application interface {
	GetAllReceipts() ([]entity.Receipt, error)
	GetReceipt(id int) (*entity.Receipt, error)
	GetAllIngredients(title string) ([]entity.Ingredient, error)
	GetIngredient(id int) (*entity.Ingredient, error)
}
