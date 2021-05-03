package application

import (
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
)

type App struct {
	storage repository.Storage
}

func NewApp(storage repository.Storage) *App {
	return &App{storage}
}

func (app *App) GetAllReceipts() ([]entity.Receipt, error) {
	return app.storage.GetReceipts()
}

func (app *App) GetReceipt(id int) (*entity.Receipt, error) {
	return app.storage.GetReceipt(id)
}

func (app *App) GetAllIngredients(title string) ([]entity.Ingredient, error) {
	return app.storage.GetIngredients(title)
}

func (app *App) GetIngredient(id int) (*entity.Ingredient, error) {
	return app.storage.GetIngredient(id)
}
