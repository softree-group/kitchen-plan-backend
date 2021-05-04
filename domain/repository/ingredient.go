package repository

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type IngredientReceiver interface {
	Filter(title string) ([]entity.Ingredient, error)
	Receive(id int) (*entity.Ingredient, error)
	ForReceipt(id int) ([]entity.Ingredient, error)
}
