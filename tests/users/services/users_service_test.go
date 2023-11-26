package users_service_test

import (
	model_interfaces "match-system/interfaces/models"
	"match-system/internal/store"
	"match-system/plugins"
	"testing"

	user_models "match-system/internal/user/models"
	user_services "match-system/internal/user/services"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UsersServiceTestSuite struct {
	suite.Suite
	logger *plugins.Logger
	store  *store.Memory

	service       *user_services.UsersService
	existingUser1 *user_models.User
	existingUser2 *user_models.User
}

func (suite *UsersServiceTestSuite) SetupTest() {
	env := plugins.NewEnv()
	suite.logger = plugins.NewLogger(env)
	suite.store = store.MemoryStore

	service := user_services.NewUsersService(suite.logger)
	suite.service, _ = service.(*user_services.UsersService)
}

func (suite *UsersServiceTestSuite) SetMultiExitedUserToStore() {
	suite.existingUser1 = &user_models.User{
		ID:          "user_id_1",
		Name:        "ExistingUser1",
		Height:      170,
		Gender:      "female",
		WantedDates: 2,
	}

	suite.existingUser2 = &user_models.User{
		ID:          "user_id_2",
		Name:        "ExistingUser2",
		Height:      171,
		Gender:      "female",
		WantedDates: 3,
	}

	suite.store.SetUsers([]model_interfaces.User{suite.existingUser1, suite.existingUser2})
}

func (suite *UsersServiceTestSuite) SetZeroWantedDateMultiExitedUserToStore() {
	suite.existingUser1 = &user_models.User{
		ID:          "user_id_1",
		Name:        "ExistingUser1",
		Height:      170,
		Gender:      "female",
		WantedDates: 0,
	}

	suite.existingUser2 = &user_models.User{
		ID:          "user_id_2",
		Name:        "ExistingUser2",
		Height:      171,
		Gender:      "female",
		WantedDates: 0,
	}

	suite.store.SetUsers([]model_interfaces.User{suite.existingUser1, suite.existingUser2})
}

func (suite *UsersServiceTestSuite) TestUsersService_AddUserAndMatch_Sucess() {
	suite.SetMultiExitedUserToStore()
	newUser := &user_models.User{
		Name:        "TestUser",
		Height:      180,
		Gender:      "male",
		WantedDates: 3,
	}

	result, err := suite.service.AddUserAndMatch(newUser)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), 2, len(result.GetMatches()))
	assert.Equal(suite.T(), "user_id_1", result.GetMatches()[0].GetID())
	assert.Equal(suite.T(), "user_id_2", result.GetMatches()[1].GetID())
}

func (suite *UsersServiceTestSuite) TestUsersService_AddUserAndMatch_ZeroMatch() {
	suite.SetMultiExitedUserToStore()
	newUser := &user_models.User{
		ID:          uuid.New().String(),
		Name:        "TestUser",
		Height:      160,
		Gender:      "male",
		WantedDates: 5,
	}

	result, err := suite.service.AddUserAndMatch(newUser)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), 0, len(result.GetMatches()))
}

func (suite *UsersServiceTestSuite) TestUsersService_AddUserAndMatch_WhenDatedCountZero() {
	suite.SetMultiExitedUserToStore()
	newUser := &user_models.User{
		ID:          uuid.New().String(),
		Name:        "TestUser",
		Height:      180,
		Gender:      "male",
		WantedDates: 0,
	}

	result, err := suite.service.AddUserAndMatch(newUser)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), 0, len(result.GetMatches()))
}

func (suite *UsersServiceTestSuite) TestUsersService_RemoveUser_Sucess() {
	suite.SetMultiExitedUserToStore()

	err := suite.service.RemoveTargetUser(suite.existingUser1.GetID())
	assert.NoError(suite.T(), err)
	assert.NotContains(suite.T(), suite.store.GetUsers(), suite.existingUser1)
}

func (suite *UsersServiceTestSuite) TestUsersService_RemoveUser_UserNotFound() {
	err := suite.service.RemoveTargetUser(suite.existingUser1.GetID())
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "cannot find user by userId")
}

func (suite *UsersServiceTestSuite) TestUsersService_QuerySingleUsers_Sucess() {
	suite.SetMultiExitedUserToStore()

	result, err := suite.service.QuerySingleUsers(2)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 2)
	for _, user := range result {
		assert.True(suite.T(), user.GetWantedDates() > 0)
	}
}

func (suite *UsersServiceTestSuite) TestUsersService_QuerySingleUsers_NoUserWantedDate() {
	suite.SetZeroWantedDateMultiExitedUserToStore()

	result, err := suite.service.QuerySingleUsers(2)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 0)
}

func TestUsersServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UsersServiceTestSuite))
}
