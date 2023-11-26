package user_models

import (
	model_interfaces "match-system/interfaces/models"
	"reflect"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

// https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance
var _ model_interfaces.User = (*User)(nil)

type User struct {
	ID          string // uuid, generate by application
	Name        string `json:"name"`
	Height      int    `json:"height"`
	Gender      string `json:"gender"`
	WantedDates int    `json:"wanted_dates"`

	Matches []model_interfaces.User
}

func (u *User) GetID() string {
	return u.ID
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

func (u *User) GenerateID() {
	u.ID = uuid.New().String()
}

func (u *User) AddMatches(user model_interfaces.User) {
	u.Matches = append(u.Matches, user)
}

func (u *User) RemoveMatchUserRelationByID(userId string) {
	_, index, ok := lo.FindLastIndexOf(u.GetMatches(), func(match model_interfaces.User) bool {
		return match.GetID() == userId
	})

	if ok {
		u.Matches = append(u.Matches[:index], u.Matches[index+1:]...)
	}
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
