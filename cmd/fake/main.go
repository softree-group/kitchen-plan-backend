package main

import "github.com/softree-group/kitchen-plan-backend/config"

func init() {
	config.InitConfig()
}

func main() {
	serveAPI()
}
