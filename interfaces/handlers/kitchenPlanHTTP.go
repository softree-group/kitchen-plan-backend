package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/application"
)

type KitchenPlanHTTP struct {
	kitchenPlan application.KitchenPlanAppInterface
	logger      *logrus.Logger
}

func NewKitchenPlanHTTP(kitchenPlan application.KitchenPlanAppInterface) *KitchenPlanHTTP {
	return &KitchenPlanHTTP{kitchenPlan, logrus.New()}
}
