package application

import (
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
)

type Ingredients struct {
	reps *repository.Repositories
}

func (i Ingredients) Filter(title string) ([]entity.Ingredient, error) {
	panic("implement me")
}

func (i Ingredients) Receive(id int) (*entity.Ingredient, error) {
	panic("implement me")
}

var _ IngredientReceiver = Ingredients{}

func NewIngredients(reps *repository.Repositories) *Ingredients {
	return &Ingredients{reps}
}
