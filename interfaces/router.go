package interfaces

import (
	"github.com/fasthttp/router"
	"github.com/softree-group/kitchen-plan-backend/application"
	"github.com/softree-group/kitchen-plan-backend/infrastructure/persistance"
	"github.com/softree-group/kitchen-plan-backend/interfaces/handlers"
	"net/http"
)

func NewRouter() *router.Router {
	kitchenPlanApp := application.NewKitchenPlanApp(persistance.NewKitchenPlanPGS())
	handler := handlers.NewKitchenPlanHTTP(kitchenPlanApp)

	r := router.New()
	r.Handle(http.MethodGet, "/receipts", handler.GetAllReceipts)
	r.Handle(http.MethodGet, "/receipts/{id}", handler.GetReceipt)

	r.Handle(http.MethodGet, "/ingredients", handler.GetAllIngredients)
	r.Handle(http.MethodGet, "/ingredients/{id}", handler.GetIngredient)

	return r
}
