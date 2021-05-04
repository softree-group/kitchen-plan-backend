package main

import "github.com/softree-group/kitchen-plan-backend/src/config"

func init() {
	config.InitConfig()
}

func main() {
	serveAPI()
}
