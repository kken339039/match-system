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

func (us *UsersService) AddSinglePersonAndMatch(newUser model_interfaces.User) error {
	us.store.Mutex.Lock()
	defer us.store.Mutex.Unlock()

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
			us.store.SetMatched(append(us.store.GetMatched(), newUser, user))

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

	return nil
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

	for i, mu := range us.store.GetMatched() {
		if user.IsSameUser(mu) {
			index = i
			break
		}
	}

	matched := us.store.GetMatched()
	if index >= 0 {
		us.store.SetMatched(append(matched[:index], matched[index+1:]...))
	}
}
