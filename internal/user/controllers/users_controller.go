package users_controllers

import (
	"encoding/json"
	"fmt"
	"match-system/plugin"
	"net/http"

	users_models "match-system/internal/user/models"
	users_services "match-system/internal/user/services"

	"github.com/gorilla/mux"
)

type UsersController struct {
	logger  *plugin.Logger
	env     *plugin.Env
	service *users_services.UsersService
}

func RegisterController(router *mux.Router, logger *plugin.Logger, env *plugin.Env) {
	service := users_services.NewUsersService(logger)
	uc := &UsersController{
		logger:  logger,
		env:     env,
		service: service,
	}
	router.HandleFunc("/api/AddSinglePersonAndMatch", uc.AddSinglePersonAndMatch).Methods("POST")
}

func (uc *UsersController) AddSinglePersonAndMatch(w http.ResponseWriter, r *http.Request) {
	var newInstance users_models.User
	err := json.NewDecoder(r.Body).Decode(&newInstance)
	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to decode user, error: %s", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := uc.service.AddSinglePersonAndMatch(&newInstance)
	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed create user to match, error: %s", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to encode newUser, error: %s", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
