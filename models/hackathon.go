package models

import "time"

type Hackathon struct {
	Id                   int
	Name                 string
	Description          string
	Location             string
	StartDate            time.Time
	TeamSize             int
	MaxNumberOfTeams     int
	OpenForSubscriptions bool
	Teams                []Team
}
