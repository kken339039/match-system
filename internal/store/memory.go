package store

import (
	model_interfaces "match-system/interfaces/models"
	store_interfaces "match-system/interfaces/stores"
	"sync"
)

var MemoryStore *Memory

func init() {
	MemoryStore = &Memory{}
}

// https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance
var _ store_interfaces.Memory = (*Memory)(nil)

type Memory struct {
	Users []model_interfaces.User

	Mutex sync.Mutex
}

func (m *Memory) GetUsers() []model_interfaces.User {
	return m.Users
}

func (m *Memory) SetUsers(users []model_interfaces.User) {
	m.Users = users
}
