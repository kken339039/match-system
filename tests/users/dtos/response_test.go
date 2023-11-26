package users_response_test

import (
	model_interfaces "match-system/interfaces/models"
	"match-system/internal/user/dtos"
	user_models "match-system/internal/user/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAddSinglePersonAndMatchResponse(t *testing.T) {
	matchedUser1 := &user_models.User{
		ID:          "user_id_002",
		Name:        "user_name_2",
		Height:      170,
		Gender:      "female",
		WantedDates: 2,
	}
	matchedUser2 := &user_models.User{
		ID:          "user_id_003",
		Name:        "user_name_3",
		Height:      171,
		Gender:      "female",
		WantedDates: 1,
	}
	matchedUser3 := &user_models.User{
		ID:          "user_id_004",
		Name:        "user_name_5",
		Height:      172,
		Gender:      "female",
		WantedDates: 3,
	}

	userWithMatches := &user_models.User{
		ID:          "user_id_001",
		Name:        "user_name_1",
		Height:      180,
		Gender:      "male",
		WantedDates: 3,
		Matches:     []model_interfaces.User{matchedUser1, matchedUser2, matchedUser3},
	}

	response := dtos.ParseAddSinglePersonAndMatchResponse(userWithMatches)

	assert.NotNil(t, response)
	assert.Equal(t, userWithMatches.GetID(), response.ID)
	assert.Equal(t, userWithMatches.GetName(), response.Name)
	assert.Equal(t, userWithMatches.GetHeight(), response.Height)
	assert.Equal(t, userWithMatches.GetGender(), response.Gender)
	assert.Equal(t, userWithMatches.GetWantedDates(), response.WantedDates)
	assert.Len(t, response.Matches, len(userWithMatches.GetMatches()))

	for i, match := range userWithMatches.GetMatches() {
		assert.Equal(t, match.GetID(), response.Matches[i].ID)
		assert.Equal(t, match.GetName(), response.Matches[i].Name)
		assert.Equal(t, match.GetHeight(), response.Matches[i].Height)
		assert.Equal(t, match.GetGender(), response.Matches[i].Gender)
	}
}

func TestParseQuerySinglePeopleResponse(t *testing.T) {
	// 準備測試數據，包括多個具有不同屬性的用戶
	user1 := &user_models.User{
		ID:          "user_id_001",
		Name:        "user_name_1",
		Height:      170,
		Gender:      "female",
		WantedDates: 2,
	}
	user2 := &user_models.User{
		ID:          "user_id_002",
		Name:        "user_name_2",
		Height:      171,
		Gender:      "female",
		WantedDates: 1,
	}
	user3 := &user_models.User{
		ID:          "user_id_003",
		Name:        "user_name_3",
		Height:      172,
		Gender:      "female",
		WantedDates: 3,
	}
	user4 := &user_models.User{
		ID:          "user_id_004",
		Name:        "user_name_4",
		Height:      182,
		Gender:      "male",
		WantedDates: 1,
	}
	users := []model_interfaces.User{user1, user2, user3, user4}

	response := dtos.ParseQuerySinglePeopleResponse(users)

	assert.NotNil(t, response)
	assert.Len(t, response.People, len(users))
	for i, user := range users {
		assert.Equal(t, user.GetID(), response.People[i].ID)
		assert.Equal(t, user.GetName(), response.People[i].Name)
		assert.Equal(t, user.GetHeight(), response.People[i].Height)
		assert.Equal(t, user.GetGender(), response.People[i].Gender)
		assert.Equal(t, user.GetWantedDates(), response.People[i].WantedDates)
	}
}
