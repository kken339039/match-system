package main

import (
	"match-system/internal"
	"match-system/plugin"
)

func main() {
	env := plugin.NewEnv()
	env.SetDefaultEnv(map[string]string{
		"ENVIRONMENT": "development",
		"PORT":        "3000",
	})
	plugin.NewLogger(env)

	mux := internal.NewMux()
	mux.Serve()
}
