package user_models

import (
	model_interfaces "match-system/interfaces/models"
)

// https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance
var _ model_interfaces.User = (*User)(nil)

type User struct {
	Name        string `json:"name"`
	Height      int    `json:"height"`
	Gender      string `json:"gender"`
	WantedDates int    `json:"wanted_dates"`

	Matches []model_interfaces.User
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetHeight() int {
	return u.Height
}

func (u *User) GetGender() string {
	return u.Gender
}

func (u *User) GetWantedDates() int {
	return u.WantedDates
}

func (u *User) GetMatches() []model_interfaces.User {
	return u.Matches
}
