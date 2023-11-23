package plugins

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var SysEnv *Env

type Env struct {
	defaultValues map[string]string
}

func NewEnv() *Env {
	env := &Env{
		defaultValues: map[string]string{
			"ENVIRONMENT": "development",
		},
	}

	projectRoot, _ := os.Getwd()
	path := fmt.Sprintf("%s/.env", projectRoot)

	if err := godotenv.Load(path); err != nil {
		fmt.Printf("No .env file found %s", err)
	}

	SysEnv = env
	return env
}

func (e *Env) SetDefaultEnv(values map[string]string) {
	for key, value := range values {
		e.defaultValues[key] = value
	}
}

func (e *Env) GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return e.defaultValues[key]
}
