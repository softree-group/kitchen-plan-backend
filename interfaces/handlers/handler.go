package handlers

import (
	"github.com/softree-group/kitchen-plan-backend/application"
)

type Handler struct {
	kitchenPlan application.AppDelegator
}

func NewKitchenPlanHTTP(kitchenPlan application.AppDelegator) *Handler {
	return &Handler{kitchenPlan}
}
