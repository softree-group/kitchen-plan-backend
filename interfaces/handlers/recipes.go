package handlers

import (
	"encoding/json"
	"github.com/softree-group/kitchen-plan-backend/domain/entity"
	"github.com/valyala/fasthttp"
	"strconv"
)

func (handler *Handler) FilterRecipes(ctx *fasthttp.RequestCtx) {
	var filter entity.ReceiptFilter
	since := string(ctx.QueryArgs().Peek("since"))
	if since != "" {
		sinceInt, err := strconv.Atoi(since)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusBadRequest)
			return
		}
		filter.Since = sinceInt
	}

	limit := string(ctx.QueryArgs().Peek("limit"))
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusBadRequest)
			return
		}
		filter.Limit = limitInt
	}

	filter.Title = string(ctx.QueryArgs().Peek("title"))
	filter.Type = string(ctx.QueryArgs().Peek("type"))
	ingredients := ctx.QueryArgs().PeekMulti("ingredients")
	if len(ingredients) > 0 {
		ingredientsInt, err := convertBytesToInt(ingredients)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusBadRequest)
			return
		}
		filter.Ingredients = ingredientsInt
	}

	recipes, err := handler.app.Recipes.Filter(&filter)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(recipes)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetBody(body)
	ctx.SetContentType("application/json")
}

func convertBytesToInt(bytes [][]byte) ([]int, error) {
	var result []int
	for _, strId := range bytes {
		idInt, err := strconv.Atoi(string(strId))
		if err != nil {
			return nil, err
		}
		result = append(result, idInt)
	}
	return result, nil
}

func (handler *Handler) ReceiveReceipt(ctx *fasthttp.RequestCtx) {
	receiptId, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	receipt, err := handler.app.Recipes.Receive(receiptId)

	if err != nil {
		if err == entity.ErrNotFound {
			ctx.Error(err.Error(), fasthttp.StatusNotFound)
			return
		}
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(receipt)

	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

	ctx.SetBody(body)
	ctx.SetContentType("application/json")
}
