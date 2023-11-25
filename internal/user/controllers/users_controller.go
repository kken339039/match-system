package users_controllers

import (
	"encoding/json"
	"fmt"
	"match-system/plugins"
	"net/http"
	"strconv"

	service_interfaces "match-system/interfaces/services"
	"match-system/internal/user/dtos"
	user_models "match-system/internal/user/models"
	users_services "match-system/internal/user/services"
	"match-system/plugins/http_server"

	"github.com/gorilla/mux"
)

type UsersController struct {
	logger  *plugins.Logger
	env     *plugins.Env
	service service_interfaces.UsersService
}

func RegisterController(router *mux.Router, logger *plugins.Logger, env *plugins.Env) {
	service := users_services.NewUsersService(logger)
	uc := New(logger, env, service)

	router.HandleFunc("/api/users", uc.AddSinglePersonAndMatch).Methods("POST")
	router.HandleFunc("/api/users/query_single", uc.QuerySinglePeople).Methods("GET")
	router.HandleFunc("/api/users/{userId}", uc.RemoveSinglePerson).Methods("DELETE")
}

func New(logger *plugins.Logger, env *plugins.Env, service service_interfaces.UsersService) *UsersController {
	return &UsersController{
		logger:  logger,
		env:     env,
		service: service,
	}
}

// @Summary Add a new user and find matches
// @Description Add a new user to the matching system and find any possible matches for the new user.
// @ID add-single-person-and-match
// @Accept  json
// @Produce  json
// @Param input body dtos.AddSinglePersonAndMatchRequest true "New user details"
// @Success 200 {object} dtos.AddSinglePersonAndMatchResponse
// @Router /api/users [post]
func (uc *UsersController) AddSinglePersonAndMatch(w http.ResponseWriter, r *http.Request) {
	var dto dtos.AddSinglePersonAndMatchRequest
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to decode user, error: %s", err))
		http_server.BadRequestError(w, r, err)
		return
	}

	newUser := &user_models.User{
		Name:        dto.Name,
		Height:      dto.Height,
		Gender:      dto.Gender,
		WantedDates: dto.WantedDates,
	}

	user, err := uc.service.AddUserAndMatch(newUser)
	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed create user to match, error: %s", err))
		http_server.InternalServerError(w, r, err)
		return
	}

	http_server.Resoponse(w, r, dtos.ParseAddSinglePersonAndMatchResponse(user))
}

func (uc *UsersController) RemoveSinglePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dto := dtos.RemoveSinglePersonRequest{
		UserID: vars["userId"],
	}
	err := uc.service.RemoveTargetUser(dto.UserID)

	if err != nil {
		uc.logger.Error(fmt.Sprintf("Failed to remove target user, error: %s", err))
		http_server.BadRequestError(w, r, err)
		return
	}

	http_server.EmptyResoponse(w, r)
}

func (uc *UsersController) QuerySinglePeople(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	var err error
	var qeuryCount int
	if queryParams.Get("N") == "" {
		qeuryCount = 1
	} else {
		qeuryCount, err = strconv.Atoi(queryParams.Get("N"))
	}

	if err != nil {
		uc.logger.Error(fmt.Sprintf("Cannot not parse query count N, err: %s", err))
		http_server.BadRequestError(w, r, err)
		return
	}

	dto := dtos.QuerySinglePeopleRequest{
		QueryCount: qeuryCount,
	}

	result, err := uc.service.QuerySingleUsers(dto.QueryCount)

	if err != nil {
		uc.logger.Error(fmt.Sprintf("Faile to Query Single Users, err: %s", err))
		http_server.InternalServerError(w, r, err)
		return
	}

	http_server.Resoponse(w, r, dtos.ParseQuerySinglePeopleResponse(result))
}
