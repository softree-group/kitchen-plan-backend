package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/mark-by/logutils"
	"github.com/sirupsen/logrus"
	"github.com/softree-group/kitchen-plan-backend/config"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

func serveAPI() {
	r := router.New()

	startServer(r)
}

func startServer(r *router.Router) {
	logrus.Infof("Start fake server on %s:%s...", viper.GetString(config.IP), viper.GetString(config.Port))

	err := fasthttp.ListenAndServe(fmt.Sprintf("%s:%s", viper.GetString(config.IP), viper.GetString(config.Port)),
		logutils.DebugMiddleWare(r.Handler))

	if err != nil {
		logrus.Fatalf("Fail start server with error: %s", err)
	}
}
