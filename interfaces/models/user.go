package model_interfaces

type User interface {
	GetID() string
	GetName() string
	GetHeight() int
	GetGender() string
	GetWantedDates() int
	GetMatches() []User

	GenerateID()
	AddMatches(User)
	RemoveMatchUserRelationByID(string)
	DecreaseDateCount()
	IsSameUser(User) bool
}
