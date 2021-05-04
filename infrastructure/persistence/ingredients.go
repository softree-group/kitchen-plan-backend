package persistence

import (
	"github.com/jackc/pgx"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
)

type IngredientsReceiver struct {
	db *pgx.ConnPool
}

func (i IngredientsReceiver) GetIngredients(title string) ([]entity.Ingredient, error) {
	panic("implement me")
}

func (i IngredientsReceiver) GetIngredient(id int) (*entity.Ingredient, error) {
	panic("implement me")
}

func NewIngredientsReceiver(db *pgx.ConnPool) *IngredientsReceiver {
	return &IngredientsReceiver{db}
}

var _ repository.IngredientReceiver = &IngredientsReceiver{}
