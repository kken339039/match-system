package users_services

import (
	// "match-system/internal/store"

	model_interfaces "match-system/interfaces/models"
	store "match-system/internal/store"
	"match-system/plugin"
)

type UsersService struct {
	logger *plugin.Logger
	store  *store.Memory
}

func NewUsersService(logger *plugin.Logger) *UsersService {
	return &UsersService{
		logger: logger,
		store:  store.MemoryStore,
	}
}

func (us *UsersService) AddSinglePersonAndMatch(newUser model_interfaces.User) (model_interfaces.User, error) {
	us.store.Mutex.Lock()
	defer us.store.Mutex.Unlock()

	newUser.GenerateID()
	allUsers := append(us.store.GetUsers(), newUser)
	us.store.SetUsers(allUsers)

	for _, user := range allUsers {
		if (user.GetGender() != newUser.GetGender()) &&
			(newUser.GetGender() == "male" && newUser.GetHeight() > user.GetHeight()) ||
			(newUser.GetGender() == "female" && newUser.GetHeight() < user.GetHeight()) {

			if newUser.GetWantedDates() <= 0 || user.GetWantedDates() <= 0 {
				continue
			}

			newUser.AddMatches(user)
			user.AddMatches(newUser)

			newUser.DecreaseDateCount()
			user.DecreaseDateCount()

			if newUser.GetWantedDates() <= 0 {
				us.removeUser(newUser)
			}

			if user.GetWantedDates() <= 0 {
				us.removeUser(user)
			}
			break
		}
	}

	return newUser, nil
}

func (us *UsersService) removeUser(user model_interfaces.User) {
	var index int

	for i, u := range us.store.GetUsers() {
		if user.IsSameUser(u) {
			index = i
			break
		}
	}

	users := us.store.GetUsers()
	if index >= 0 {
		us.store.SetUsers(append(users[:index], users[index+1:]...))
	}
}
