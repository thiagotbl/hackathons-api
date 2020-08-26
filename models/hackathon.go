package models

import (
	"errors"
	"strings"
	"time"
)

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

func (hackathon *Hackathon) Validate() error {
	if strings.TrimSpace(hackathon.Name) == "" {
		return errors.New("name should not be empty")
	}
	if strings.TrimSpace(hackathon.Location) == "" {
		return errors.New("location should not be empty")
	}
	if hackathon.StartDate.IsZero() {
		return errors.New("invalid start date")
	}
	if hackathon.TeamSize < 1 {
		return errors.New("team size should be greater than zero")
	}
	if hackathon.NumberOfTeams < 1 {
		return errors.New("number of teams should be greater than zero")
	}

	return nil
}
