package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (handler *Handler) FilterIngredients(ctx *fasthttp.RequestCtx) {
	title := string(ctx.QueryArgs().Peek("title"))
	if title == "" {
		ctx.Error("No title specified", fasthttp.StatusBadRequest)
		return
	}

	ingredients, err := handler.app.Ingredients.Filter(title)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(ingredients)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetBody(body)
	ctx.SetContentType("application/json")
}

func (handler *Handler) ReceiveIngredient(ctx *fasthttp.RequestCtx) {

}
