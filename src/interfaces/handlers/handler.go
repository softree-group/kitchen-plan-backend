package handlers

import (
	"github.com/softree-group/kitchen-plan-backend/src/application"
)

type Handler struct {
	app *application.Application
}

func NewHandler(app *application.Application) *Handler {
	return &Handler{app}
}
