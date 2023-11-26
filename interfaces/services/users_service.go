package service_interfaces

import model_interfaces "match-system/interfaces/models"

type UsersService interface {
	AddUserAndMatch(newuser model_interfaces.User) (model_interfaces.User, error)
	RemoveTargetUser(userId string) error
	QuerySingleUsers(quertCount int) ([]model_interfaces.User, error)
}
