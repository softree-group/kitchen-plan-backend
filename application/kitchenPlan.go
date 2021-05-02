package application

import (
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
)

type KitchenPlanApp struct {
	receiptsDB repository.KitchenPlanDBRepo
}

type KitchenPlanAppInterface interface {
	GetAllReceipts() ([]entity.Receipt, error)
	GetReceipt(id int) (*entity.Receipt, error)
	GetAllIngredients(title string) ([]entity.Ingredient, error)
	GetIngredient(id int) (*entity.Ingredient, error)
}

func NewKitchenPlanApp(receiptsDB repository.KitchenPlanDBRepo) *KitchenPlanApp {
	return &KitchenPlanApp{receiptsDB}
}

func (app *KitchenPlanApp) GetAllReceipts() ([]entity.Receipt, error) {
	return app.receiptsDB.GetReceipts()
}

func (app *KitchenPlanApp) GetReceipt(id int) (*entity.Receipt, error) {
	return app.receiptsDB.GetReceipt(id)
}

func (app *KitchenPlanApp) GetAllIngredients(title string) ([]entity.Ingredient, error) {
	return app.receiptsDB.GetIngredients(title)
}

func (app *KitchenPlanApp) GetIngredient(id int) (*entity.Ingredient, error) {
	return app.receiptsDB.GetIngredient(id)
}
