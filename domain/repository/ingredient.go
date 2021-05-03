package repository

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type IngredientReceiver interface {
	GetIngredients(title string) ([]entity.Ingredient, error)
	GetIngredient(id int) (*entity.Ingredient, error)
}
