package application

import (
	"github.com/softree-group/kitchen-plan-backend/src/config"
	"github.com/softree-group/kitchen-plan-backend/src/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/src/domain/repository"
	"github.com/spf13/viper"
	"math/rand"
)

type Ingredients struct {
	reps *repository.Repositories
}

func (i Ingredients) Filter(title string) ([]entity.Ingredient, error) {
	ingredients, err := i.reps.IngredientReceiver.Filter(title)
	if err != nil {
		return nil, err
	}

	root := viper.GetString(config.StaticStorageRoot)
	for idx := 0; idx < len(ingredients); idx++ {
		ingredients[idx].SetImageRoot(root)
	}
	return ingredients, nil
}

func (i Ingredients) Receive(id int) (*entity.Ingredient, error) {
	ingredient, err := i.reps.IngredientReceiver.Receive(id)
	if err != nil {
		return nil, err
	}

	ingredient.Proteins = rand.Intn(40)
	ingredient.Fats = rand.Intn(40)
	ingredient.Carbohydrates = rand.Intn(40)

	ingredient.SetImageRoot(viper.GetString(config.StaticStorageRoot))

	return ingredient, nil
}

var _ IngredientReceiver = Ingredients{}

func NewIngredients(reps *repository.Repositories) *Ingredients {
	return &Ingredients{reps}
}
