package users_services

import (
	"errors"
	model_interfaces "match-system/interfaces/models"
	service_interfaces "match-system/interfaces/services"
	store "match-system/internal/store"

	"match-system/plugins"

	"github.com/samber/lo"
)

type UsersService struct {
	logger *plugins.Logger
	store  *store.Memory
}

func NewUsersService(logger *plugins.Logger) service_interfaces.UsersService {
	return &UsersService{
		logger: logger,
		store:  store.MemoryStore,
	}
}

func (us *UsersService) AddUserAndMatch(newUser model_interfaces.User) (model_interfaces.User, error) {
	us.store.Mutex.Lock()
	defer us.store.Mutex.Unlock()

	newUser.GenerateID()

	for _, user := range us.store.GetUsers() {
		if (user.GetGender() != newUser.GetGender()) &&
			(newUser.GetGender() == "male" && newUser.GetHeight() > user.GetHeight()) ||
			(newUser.GetGender() == "female" && newUser.GetHeight() < user.GetHeight()) {

			newUser.AddMatches(user)
			user.AddMatches(newUser)

			newUser.DecreaseDateCount()
			user.DecreaseDateCount()

			var err error
			if newUser.GetWantedDates() <= 0 {
				err = removeUser(us.store, newUser.GetID())
			}

			if user.GetWantedDates() <= 0 {
				err = removeUser(us.store, newUser.GetID())
			}

			if err != nil {
				return nil, err
			}
		}
	}

	allUsers := append(us.store.GetUsers(), newUser)
	us.store.SetUsers(allUsers)
	return newUser, nil
}

func (us *UsersService) RemoveTargetUser(userId string) error {
	us.store.Mutex.Lock()
	defer us.store.Mutex.Unlock()

	err := removeUser(us.store, userId)

	if err != nil {
		return err
	}

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

func removeUser(store *store.Memory, userId string) error {
	users := store.GetUsers()
	user, index, ok := lo.FindLastIndexOf(users, func(u model_interfaces.User) bool {
		return u.GetID() == userId
	})

	if ok {
		store.SetUsers(append(users[:index], users[index+1:]...))
		return nil
	}

	for _, matchedUser := range user.GetMatches() {
		matchedUser.RemoveMatchUserRelationByID(userId)
	}

	return errors.New("cannot find user by userId")
}
