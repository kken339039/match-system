package users_services

import (
	"errors"
	"fmt"
	model_interfaces "match-system/interfaces/models"
	store "match-system/internal/store"

	"match-system/plugins"

	"github.com/samber/lo"
)

type UsersService struct {
	logger *plugins.Logger
	store  *store.Memory
}

func NewUsersService(logger *plugins.Logger) *UsersService {
	return &UsersService{
		logger: logger,
		store:  store.MemoryStore,
	}
}

func (us *UsersService) AddUserAndMatch(newUser model_interfaces.User) (model_interfaces.User, error) {
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
		}
	}

	return newUser, nil
}

func (us *UsersService) RemoveTargetUser(userId string) error {
	us.store.Mutex.Lock()
	defer us.store.Mutex.Unlock()

	allUsers := us.store.GetUsers()
	_, targetIndex, ok := lo.FindLastIndexOf(allUsers, func(user model_interfaces.User) bool {
		return user.GetID() == userId
	})

	if !ok {
		errorMsg := fmt.Sprintf("User Not found, userId: %s", userId)
		us.logger.Error(errorMsg)
		return errors.New(errorMsg)
	}

	us.store.SetUsers(append(allUsers[:targetIndex], allUsers[targetIndex+1:]...))
	return nil
}

func (us *UsersService) QuerySingleUsers(queryCount int) ([]model_interfaces.User, error) {
	us.store.Mutex.Lock()
	defer us.store.Mutex.Unlock()

	var results []model_interfaces.User
	for _, u := range us.store.GetUsers() {
		if u.GetWantedDates() > 0 {
			results = append(results, u)
			if len(results) == queryCount {
				break
			}
		}
	}
	return results, nil
}
