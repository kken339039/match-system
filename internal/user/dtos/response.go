package dtos

import model_interfaces "match-system/interfaces/models"

type AddSinglePersonAndMatchResponse struct {
	PeopleResponse
	Matches []MatchedUserResponse
}

type MatchedUserResponse struct {
	ID     string
	Name   string
	Height int
	Gender string
}

type QuerySinglePeopleResponse struct {
	Items []PeopleResponse
}

type PeopleResponse struct {
	ID          string
	Name        string
	Height      int
	Gender      string
	WantedDates int
}

func ParseAddSinglePersonAndMatchResponse(user model_interfaces.User) *AddSinglePersonAndMatchResponse {
	matches := []MatchedUserResponse{}
	for _, match := range user.GetMatches() {
		matches = append(matches, MatchedUserResponse{
			match.GetID(),
			match.GetName(),
			match.GetHeight(),
			match.GetGender(),
		})
	}
	people := PeopleResponse{
		ID:          user.GetID(),
		Name:        user.GetName(),
		Height:      user.GetHeight(),
		Gender:      user.GetGender(),
		WantedDates: user.GetWantedDates(),
	}

	return &AddSinglePersonAndMatchResponse{
		people,
		matches,
	}
}

func ParseQuerySinglePeopleResponse(result []model_interfaces.User) *QuerySinglePeopleResponse {
	peoples := []PeopleResponse{}
	for _, user := range result {
		peoples = append(peoples, PeopleResponse{
			ID:          user.GetID(),
			Name:        user.GetName(),
			Height:      user.GetHeight(),
			Gender:      user.GetGender(),
			WantedDates: user.GetWantedDates(),
		})
	}

	return &QuerySinglePeopleResponse{
		Items: peoples,
	}
}
