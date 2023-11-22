package repository

import (
	model_interfaces "match-system/interfaces/models"
	repository_interfaces "match-system/interfaces/repositories"
	"sync"
)

// https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance
var _ repository_interfaces.Store = (*Store)(nil)

type Store struct {
	Users   []model_interfaces.User
	Matched []model_interfaces.User

	Mutex sync.Mutex
}

func (s *Store) GetUsers() []model_interfaces.User {
	return s.Users
}

func (s *Store) GetMatched() []model_interfaces.User {
	return s.Matched
}
