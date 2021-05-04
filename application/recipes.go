package application

import (
	"github.com/softree-group/kitchen-plan-backend/config"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
	"github.com/spf13/viper"
)

type Recipes struct {
	reps *repository.Repositories
}

func (r Recipes) Filter(filter entity.ReceiptFilter) ([]entity.Receipt, error) {
	panic("implement me")
}

func (r Recipes) Receive(id int) (*entity.Receipt, error) {
	receipt, err := r.reps.ReceiptReceiver.Receive(id)
	if err != nil {
		return nil, err
	}
	ingredients, err := r.reps.IngredientReceiver.ForReceipt(id)
	if err != nil {
		return nil, err
	}

	receipt.Ingredients = ingredients

	receipt.SetImageRoot(viper.GetString(config.StaticStorageRoot))
	return receipt, nil
}

var _ ReceiptReceiver = Recipes{}

func NewRecipes(reps *repository.Repositories) *Recipes {
	return &Recipes{reps}
}
