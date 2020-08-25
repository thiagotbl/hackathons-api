package models

import "time"

type Hackathon struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Location          string    `json:"location"`
	StartDate         time.Time `json:"startDate"`
	TeamSize          int       `json:"teamSize"`
	NumberOfTeams     int       `json:"numberOfTeams"`
	SubscriptionsOpen bool      `json:"subscriptions"`
	Teams             []Team    `json:"teams"`
}
