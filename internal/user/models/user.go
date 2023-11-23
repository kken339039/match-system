package user_models

import (
	model_interfaces "match-system/interfaces/models"
	"reflect"

	"github.com/samber/lo"
)

// https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance
var _ model_interfaces.User = (*User)(nil)

type User struct {
	Name        string `json:"name"`
	Height      int    `json:"height"`
	Gender      string `json:"gender"`
	WantedDates int    `json:"wanted_dates"`

	Matches []User
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
	matches := []model_interfaces.User{}
	for _, match := range u.Matches {
		matches = append(matches, lo.ToPtr(match))
	}

	return matches
}

func (u *User) AddMatches(user model_interfaces.User) {
	userInstance, _ := user.(*User)
	u.Matches = append(u.Matches, *userInstance)
}

func (u *User) DecreaseDateCount() {
	u.WantedDates--
}

func (u *User) IsSameUser(other model_interfaces.User) bool {
	if o, ok := other.(*User); ok {
		return reflect.DeepEqual(u, o)
	}

	return false
}
