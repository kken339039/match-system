package repository_interfaces

import model_interfaces "match-system/interfaces/models"

type Store interface {
	GetUsers() []model_interfaces.User
	GetMatched() []model_interfaces.User
}
