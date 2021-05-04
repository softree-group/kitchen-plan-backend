package handlers

import (
	"encoding/json"
	"github.com/softree-group/kitchen-plan-backend/src/domain/entity"
	"github.com/valyala/fasthttp"
	"strconv"
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

	if len(ingredients) == 0 {
		ctx.SetStatusCode(fasthttp.StatusNoContent)
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
	ingredientId, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	ingredient, err := handler.app.Ingredients.Receive(ingredientId)
	if err != nil {
		if err == entity.ErrNotFound {
			ctx.Error(err.Error(), fasthttp.StatusNotFound)
			return
		}
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(ingredient)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetBody(body)
	ctx.SetContentType("application/json")
}
