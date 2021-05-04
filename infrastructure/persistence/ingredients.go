package persistence

import (
	"database/sql"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
)

const (
	sqlIngredientsForReceipt = "ingredientsForReceipt"
)

type IngredientsReceiver struct {
	db *pgx.ConnPool
}

func (i IngredientsReceiver) Filter(title string) ([]entity.Ingredient, error) {
	panic("implement me")
}

func (i IngredientsReceiver) Receive(id int) (*entity.Ingredient, error) {
	panic("implement me")
}

func (i IngredientsReceiver) ForReceipt(id int) ([]entity.Ingredient, error) {
	tx, err := i.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {endTx(tx, err)}()

	rows, err := tx.Query(sqlIngredientsForReceipt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []entity.Ingredient
	for rows.Next() {
		ingredient := entity.Ingredient{}
		quantity := sql.NullFloat64{}
		measure := sql.NullString{}
		err = rows.Scan(&ingredient.Id, &ingredient.Title, &ingredient.Image, &quantity, &measure)
		if err != nil {
			return nil, err
		}
		if quantity.Valid {
			ingredient.Quantity = float32(quantity.Float64)
		}
		if measure.Valid {
			ingredient.Measure = measure.String
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (i IngredientsReceiver) Prepare() {
	if _, err := i.db.Prepare(sqlIngredientsForReceipt,
		"select i.id, i.title, i.image, ri.quantity, ri.measure from recipes_ingredients ri" +
		" join ingredients i on ri.ingredient_id = i.id" +
		" where ri.receipt_id = $1"); err != nil {
		logrus.Fatal(err)
	}
}

func NewIngredientsReceiver(db *pgx.ConnPool) *IngredientsReceiver {
	return &IngredientsReceiver{db}
}

var _ repository.IngredientReceiver = &IngredientsReceiver{}
