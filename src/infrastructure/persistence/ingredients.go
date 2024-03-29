package persistence

import (
	"database/sql"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/src/domain/entity"
	"github.com/softree-group/kitchen-plan-backend/src/domain/repository"
	"strings"
)

const (
	sqlIngredientsForReceipt = "ingredientsForReceipt"
	sqlReceiveIngredient     = "ingredientReceive"
	sqlIngredientsFilter     = "ingredientFilter"
)

type IngredientsReceiver struct {
	db *pgx.ConnPool
}

func (i IngredientsReceiver) Filter(title string) ([]entity.Ingredient, error) {
	tx, err := i.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { endTx(tx, err) }()
	searchTitle := strings.ToLower(title) + ":*" + "|" + strings.Title(title) + ":*"

	rows, err := tx.Query(sqlIngredientsFilter, searchTitle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []entity.Ingredient

	for rows.Next() {
		ingredient := entity.Ingredient{}
		image := sql.NullString{}
		err = rows.Scan(&ingredient.Id, &ingredient.Title, &image)
		if err != nil {
			return nil, err
		}
		if image.Valid {
			ingredient.Image = image.String
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (i IngredientsReceiver) Receive(id int) (*entity.Ingredient, error) {
	tx, err := i.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { endTx(tx, err) }()

	ingredient := entity.Ingredient{}

	ingredientImage := sql.NullString{}
	protein := sql.NullFloat64{}
	fat := sql.NullFloat64{}
	carbohydrate := sql.NullFloat64{}
	energy := sql.NullFloat64{}

	err = tx.QueryRow(sqlReceiveIngredient, id).Scan(&ingredient.Id, &ingredient.Title, &ingredientImage,
		&protein, &fat, &carbohydrate, &energy)

	if err != nil {
		return nil, toEntityError(err)
	}

	if protein.Valid {
		ingredient.Proteins = float32(protein.Float64)
	}

	if fat.Valid {
		ingredient.Fats = float32(fat.Float64)
	}

	if carbohydrate.Valid {
		ingredient.Carbohydrates = float32(carbohydrate.Float64)
	}

	if energy.Valid {
		ingredient.Energy = float32(energy.Float64)
	}

	if ingredientImage.Valid {
		ingredient.Image = ingredientImage.String
	}

	return &ingredient, nil
}

func (i IngredientsReceiver) ForReceipt(id int) ([]entity.Ingredient, error) {
	tx, err := i.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() { endTx(tx, err) }()

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
		image := sql.NullString{}
		err = rows.Scan(&ingredient.Id, &ingredient.Title, &image, &quantity, &measure)
		if err != nil {
			return nil, err
		}
		if quantity.Valid {
			ingredient.Quantity = float32(quantity.Float64)
		}
		if measure.Valid {
			ingredient.Measure = measure.String
		}
		if image.Valid {
			ingredient.Image = image.String
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (i IngredientsReceiver) Prepare() {
	if _, err := i.db.Prepare(sqlIngredientsForReceipt,
		"select i.id, i.title, i.image, ri.quantity, ri.measure from recipes_ingredients ri"+
			" join ingredients i on ri.ingredient_id = i.id"+
			" where ri.receipt_id = $1"); err != nil {
		logrus.Fatal(err)
	}

	if _, err := i.db.Prepare(sqlReceiveIngredient,
		"select id, title, image, protein, fat, carbohydrate, energy from ingredients where id = $1"); err != nil {
		logrus.Fatal(err)
	}

	if _, err := i.db.Prepare(sqlIngredientsFilter,
		"select id, title, image from ingredients where" +
		" to_tsvector(title) @@ to_tsquery('russian', $1) and not is_overall"); err != nil {
		logrus.Fatal(err)
	}
}

func NewIngredientsReceiver(db *pgx.ConnPool) *IngredientsReceiver {
	return &IngredientsReceiver{db}
}

var _ repository.IngredientReceiver = &IngredientsReceiver{}
