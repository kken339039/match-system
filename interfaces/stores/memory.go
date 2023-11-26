package repository_interfaces

import model_interfaces "match-system/interfaces/models"

type Memory interface {
	GetUsers() []model_interfaces.User
	SetUsers([]model_interfaces.User)
}
