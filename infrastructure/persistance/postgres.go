package persistance

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type PostgresStorage struct {

}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{}
}

func (storage *PostgresStorage) GetReceipts() ([]entity.Receipt, error) {
	return nil, nil
}

func (storage *PostgresStorage) GetReceipt(id int) (*entity.Receipt, error) {
	return nil, nil
}

func (storage *PostgresStorage) GetIngredients(title string) ([]entity.Ingredient, error) {
	return nil, nil
}

func (storage *PostgresStorage) GetIngredient(id int) (*entity.Ingredient, error) {
	return nil, nil
}
