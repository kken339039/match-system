package users_controllers

import (
	"encoding/json"
	"fmt"
	"match-system/plugins"
	"net/http"
	"strconv"

	"match-system/internal/user/dtos"
	users_models "match-system/internal/user/models"
	users_services "match-system/internal/user/services"

	"github.com/gorilla/mux"
)

type UsersController struct {
	logger  *plugins.Logger
	env     *plugins.Env
	service *users_services.UsersService
}

func RegisterController(router *mux.Router, logger *plugins.Logger, env *plugins.Env) {
	service := users_services.NewUsersService(logger)
	uc := &UsersController{
		logger:  logger,
		env:     env,
		service: service,
	}

	router.HandleFunc("/api/users", uc.AddSinglePersonAndMatch).Methods("POST")
	router.HandleFunc("/api/users/query_single", uc.QuerySinglePeople).Methods("GET")
	router.HandleFunc("/api/users/{userId}", uc.RemoveSinglePerson).Methods("DELETE")
}

func (uc *UsersController) AddSinglePersonAndMatch(w http.ResponseWriter, r *http.Request) {
	var newInstance users_models.User
	err := json.NewDecoder(r.Body).Decode(&newInstance)
	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to decode user, error: %s", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := uc.service.AddUserAndMatch(&newInstance)
	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed create user to match, error: %s", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(dtos.ParseAddSinglePersonAndMatchResponse(user))
	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to encode newUser, error: %s", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (uc *UsersController) RemoveSinglePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	err := uc.service.RemoveTargetUser(userId)

	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to remove target user, error: %s", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UsersController) QuerySinglePeople(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	qeuryCount, err := strconv.Atoi(queryParams.Get("N"))

	if err != nil {
		uc.logger.Error(fmt.Sprintf("Cannot not parse query count N, err: %s", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := uc.service.QuerySingleUsers(int(qeuryCount))

	if err != nil {
		uc.logger.Error(fmt.Sprintf("Faile to Query Single Users, err: %s", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(dtos.ParseQuerySinglePeopleResponse(result))
	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to encode query result, err: %s", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
