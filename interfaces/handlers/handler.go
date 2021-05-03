package handlers

import (
	"github.com/softree-group/kitchen-plan-backend/application"
)

type Handler struct {
	kitchenPlan application.Application
}

func NewKitchenPlanHTTP(kitchenPlan application.Application) *Handler {
	return &Handler{kitchenPlan}
}
