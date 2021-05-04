package application

import (
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
)

type Application struct {
	Ingredients IngredientReceiver
	Recipes     ReceiptReceiver
}

func New(reps *repository.Repositories) *Application {

	return &Application{
		Ingredients: NewIngredients(reps),
		Recipes:     NewRecipes(reps),
	}
}
