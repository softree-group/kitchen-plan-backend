package interfaces

import (
	"github.com/fasthttp/router"
	"github.com/softree-group/kitchen-plan-backend/application"
	"github.com/softree-group/kitchen-plan-backend/infrastructure"
	"github.com/softree-group/kitchen-plan-backend/interfaces/handlers"
)

func NewRouter() *router.Router {
	repos := infrastructure.NewRepositories()

	kitchenPlanApp := application.NewApp(repos)
	handler := handlers.NewKitchenPlanHTTP(kitchenPlanApp)

	r := router.New()
	r.GET("/receipts", handler.GetAllReceipts)
	r.GET("/receipts/{id}", handler.GetReceipt)

	r.GET("/ingredients", handler.GetAllIngredients)
	r.GET("/ingredients/{id}", handler.GetIngredient)

	return r
}
