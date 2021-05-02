package persistance

import "github.com/softree-group/kitchen-plan-backend/domain/entity"

type KitchenPlanPGS struct {

}

func NewKitchenPlanPGS() *KitchenPlanPGS {
	return &KitchenPlanPGS{}
}

func (manager *KitchenPlanPGS) GetReceipts() ([]entity.Receipt, error) {
	return nil, nil
}

func (manager *KitchenPlanPGS) GetReceipt(id int) (*entity.Receipt, error) {
	return nil, nil
}

func (manager *KitchenPlanPGS) GetIngredients(title string) ([]entity.Ingredient, error) {
	return nil, nil
}

func (manager *KitchenPlanPGS) GetIngredient(id int) (*entity.Ingredient, error) {
	return nil, nil
}
