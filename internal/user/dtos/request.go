package dtos

type AddSinglePersonAndMatchRequest struct {
	Name        string `json:"name"`
	Height      int    `json:"height"`
	Gender      string `json:"gender"`
	WantedDates int    `json:"wanted_dates"`
}

type RemoveSinglePersonRequest struct {
	UserID string
}

type QuerySinglePeopleRequest struct {
	QueryCount int
}
