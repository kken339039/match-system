package main

import (
	"match-system/internal"
	"match-system/plugins"
)

func main() {
	env := plugins.NewEnv()
	env.SetDefaultEnv(map[string]string{
		"ENVIRONMENT": "development",
		"PORT":        "3000",
	})
	plugins.NewLogger(env)

	mux := internal.NewMux()
	mux.Serve()
}
