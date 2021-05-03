package application

import (
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
)

type KitchenPlanApp struct {
	storage repository.KitchenPlanStorage
}

type KitchenPlanAppInterface interface {
	GetAllReceipts() ([]entity.Receipt, error)
	GetReceipt(id int) (*entity.Receipt, error)
	GetAllIngredients(title string) ([]entity.Ingredient, error)
	GetIngredient(id int) (*entity.Ingredient, error)
}

func NewKitchenPlanApp(receiptsDB repository.KitchenPlanStorage) *KitchenPlanApp {
	return &KitchenPlanApp{receiptsDB}
}

func (app *KitchenPlanApp) GetAllReceipts() ([]entity.Receipt, error) {
	return app.storage.GetReceipts()
}

func (app *KitchenPlanApp) GetReceipt(id int) (*entity.Receipt, error) {
	return app.storage.GetReceipt(id)
}

func (app *KitchenPlanApp) GetAllIngredients(title string) ([]entity.Ingredient, error) {
	return app.storage.GetIngredients(title)
}

func (app *KitchenPlanApp) GetIngredient(id int) (*entity.Ingredient, error) {
	return app.storage.GetIngredient(id)
}
