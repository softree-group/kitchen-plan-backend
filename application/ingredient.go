package application

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type appIngredientReceiver interface {
	GetAllIngredients(title string) ([]entity.Ingredient, error)
	GetIngredient(id int) (*entity.Ingredient, error)
}
