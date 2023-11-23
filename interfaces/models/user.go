package model_interfaces

type User interface {
	GetName() string
	GetHeight() int
	GetGender() string
	GetWantedDates() int
	GetMatches() []User
	AddMatches(User)
	DecreaseDateCount()
	IsSameUser(User) bool
}
