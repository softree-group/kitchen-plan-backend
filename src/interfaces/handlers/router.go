package handlers

import (
	"github.com/fasthttp/router"
	"github.com/softree-group/kitchen-plan-backend/src/application"
)

func NewRouter(app *application.Application) *router.Router {
	handler := NewHandler(app)

	r := router.New()
	r.GET("/recipes", handler.FilterRecipes)
	r.GET("/recipes/{id}", handler.ReceiveReceipt)

	r.GET("/ingredients", handler.FilterIngredients)
	r.GET("/ingredients/{id}", handler.ReceiveIngredient)

	r.GET("/ping", handler.GetHealtCheck)

	return r
}
