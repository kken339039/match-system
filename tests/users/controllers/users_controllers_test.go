package users_controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	model_interfaces "match-system/interfaces/models"
	"match-system/plugins"
	"net/http"
	"net/http/httptest"

	user_ctrl "match-system/internal/user/controllers"
	user_models "match-system/internal/user/models"
	"match-system/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UsersControllerTestSuite struct {
	suite.Suite
	logger *plugins.Logger
	env    *plugins.Env

	mockController *user_ctrl.UsersController
	mockService    *mocks.UsersService
	mockNewUser    *user_models.User
}

func (suite *UsersControllerTestSuite) SetupTest() {
	suite.env = plugins.NewEnv()
	suite.logger = plugins.NewLogger(suite.env)

	suite.mockService = new(mocks.UsersService)
	suite.mockNewUser = &user_models.User{
		ID:          uuid.New().String(),
		Name:        "TestUser",
		Height:      180,
		Gender:      "male",
		WantedDates: 3,
	}

	suite.mockController = user_ctrl.New(suite.logger, suite.env, suite.mockService)
}

func (suite *UsersControllerTestSuite) TestUsersController_AddSinglePersonAndMatch_Failed() {
	errMsg := "AddUserAndMatch failed"
	suite.mockService.On("AddUserAndMatch", mock.Anything).Return(nil, errors.New(errMsg))

	requestBody, err := json.Marshal(suite.mockNewUser)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/api/users", bytes.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	recorder := httptest.NewRecorder()
	suite.mockController.AddSinglePersonAndMatch(recorder, req)
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
}

func (suite *UsersControllerTestSuite) TestUsersController_AddSinglePersonAndMatch_Success() {
	suite.mockService.On("AddUserAndMatch", mock.Anything).Return(suite.mockNewUser, nil)

	requestBody, err := json.Marshal(suite.mockNewUser)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/api/users", bytes.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	recorder := httptest.NewRecorder()
	suite.mockController.AddSinglePersonAndMatch(recorder, req)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
}

func (suite *UsersControllerTestSuite) TestUsersController_RemoveSinglePerson_Success() {
	suite.mockService.On("RemoveTargetUser", mock.Anything).Return(nil)

	path := fmt.Sprintf("/api/users/%s", suite.mockNewUser.GetID())
	req, err := http.NewRequest("DELETE", path, nil)
	assert.NoError(suite.T(), err)

	recorder := httptest.NewRecorder()
	suite.mockController.RemoveSinglePerson(recorder, req)
	assert.Equal(suite.T(), http.StatusNoContent, recorder.Code)
}

func (suite *UsersControllerTestSuite) TestUsersController_RemoveSinglePerson_Failed() {
	errMsg := "RemoveSingleUser failed"
	suite.mockService.On("RemoveTargetUser", mock.Anything).Return(errors.New(errMsg))

	path := fmt.Sprintf("/api/users/%s", suite.mockNewUser.GetID())
	req, err := http.NewRequest("DELETE", path, nil)
	assert.NoError(suite.T(), err)

	recorder := httptest.NewRecorder()
	suite.mockController.RemoveSinglePerson(recorder, req)
	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
}

func (suite *UsersControllerTestSuite) TestUsersController_QuerySinglePeople_Success() {
	suite.mockService.On("QuerySingleUsers", mock.Anything).Return([]model_interfaces.User{}, nil)

	req, err := http.NewRequest("GET", "/api/users/query_single", nil)
	assert.NoError(suite.T(), err)

	recorder := httptest.NewRecorder()
	suite.mockController.QuerySinglePeople(recorder, req)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
}

func (suite *UsersControllerTestSuite) TestUsersController_QuerySinglePeople_Failed() {
	errMsg := "QuerySingleUsers failed"
	suite.mockService.On("QuerySingleUsers", mock.Anything).Return(nil, errors.New(errMsg))

	req, err := http.NewRequest("GET", "/api/users/query_single", nil)
	assert.NoError(suite.T(), err)

	recorder := httptest.NewRecorder()
	suite.mockController.QuerySinglePeople(recorder, req)
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
}

func TestUsersControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UsersControllerTestSuite))
}
