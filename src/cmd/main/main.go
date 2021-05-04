package main

import (
	"github.com/softree-group/kitchen-plan-backend/src/application"
	"github.com/softree-group/kitchen-plan-backend/src/config"
	"github.com/softree-group/kitchen-plan-backend/src/infrastructure"
	"github.com/softree-group/kitchen-plan-backend/src/interfaces"
)

func init() {
	config.InitConfig()
}

func main() {
	repositories := infrastructure.New()
	app := application.New(repositories)
	interfaces.ServeAPI(app)
}
