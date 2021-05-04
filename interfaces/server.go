package interfaces

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/mark-by/logutils"
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/application"
	"github.com/softree-group/kitchen-plan-backend/config"
	"github.com/softree-group/kitchen-plan-backend/interfaces/handlers"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

func startServer(r *router.Router) {
	logrus.Infof("Start server on %s:%s", viper.Get(config.IP), viper.Get(config.Port))

	err := fasthttp.ListenAndServe(
		fmt.Sprintf("%s:%s", viper.Get(config.IP), viper.Get(config.Port)),
		logutils.DebugMiddleWare(r.Handler))

	if err != nil {
		logrus.Fatalf("Fail to start server: %s", err)
	}
}

func ServeAPI(app *application.Application) {
	r := handlers.NewRouter(app)
	startServer(r)
}
