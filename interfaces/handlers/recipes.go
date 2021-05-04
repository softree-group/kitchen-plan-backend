package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

func (handler *Handler) FilterRecipes(ctx *fasthttp.RequestCtx) {

}

func (handler *Handler) ReceiveReceipt(ctx *fasthttp.RequestCtx) {
	receiptId, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	receipt, err := handler.app.Recipes.Receive(receiptId)

	if err != nil {
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
