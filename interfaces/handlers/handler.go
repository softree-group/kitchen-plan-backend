package handlers

import (
	"github.com/softree-group/kitchen-plan-backend/application"
)

type Handler struct {
	app *application.Application
}

func NewHandler(app *application.Application) *Handler {
	return &Handler{app}
}
