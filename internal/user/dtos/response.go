package dtos

import model_interfaces "match-system/interfaces/models"

type AddSinglePersonAndMatchResponse struct {
	ID          string
	Name        string
	Height      int
	Gender      string
	WantedDates int
	Matches     []MatchedUserResponse
}

type MatchedUserResponse struct {
	ID     string
	Name   string
	Height int
	Gender string
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
	return &AddSinglePersonAndMatchResponse{
		ID:          user.GetID(),
		Name:        user.GetName(),
		Height:      user.GetHeight(),
		Gender:      user.GetGender(),
		WantedDates: user.GetWantedDates(),
		Matches:     matches,
	}
}
