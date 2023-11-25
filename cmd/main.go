package main

import (
	user_ctrl "match-system/internal/user/controllers"
	"match-system/plugins"
	"match-system/plugins/http_server"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "match-system/docs"

	"github.com/gorilla/mux"
)

func main() {
	env := plugins.NewEnv()
	env.SetDefaultEnv(map[string]string{
		"ENVIRONMENT": "development",
		"PORT":        "3000",
	})
	plugins.NewLogger(env)

	logger := plugins.SysLogger
	router := mux.NewRouter()

	user_ctrl.RegisterController(router, logger, env)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	s := http_server.NewHttpServer(router)
	s.Serve()
}
