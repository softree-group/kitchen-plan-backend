package handlers

import (
	"github.com/valyala/fasthttp"
)

func (handler *Handler) GetHealtCheck(ctx *fasthttp.RequestCtx) {
	ctx.SetBody([]byte("OK"))
	ctx.SetContentType("text/plain")
}
