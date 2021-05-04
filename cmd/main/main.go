package main

import (
	"github.com/softree-group/kitchen-plan-backend/application"
	"github.com/softree-group/kitchen-plan-backend/config"
	"github.com/softree-group/kitchen-plan-backend/infrastructure"
	"github.com/softree-group/kitchen-plan-backend/interfaces"
)

func init() {
	config.InitConfig()
}

func main() {
	repositories := infrastructure.New()
	app := application.New(repositories)
	interfaces.ServeAPI(app)
}
