package models

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	pb "github.com/thiagotbl/hackatons-api/protos"
)

// Hackathon is the representation of a hackathon
type Hackathon struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Location          string    `json:"location"`
	StartDate         time.Time `json:"startDate"`
	TeamSize          int32     `json:"teamSize"`
	NumberOfTeams     int32     `json:"numberOfTeams"`
	SubscriptionsOpen bool      `json:"subscriptionsOpen" sql:",notnull"`
	// Teams             []Team    `json:"teams"`
}

// HackathonFromProto transforms a *protogen.Hackathon into a Hackathon struct
func HackathonFromProto(proto *pb.Hackathon) *Hackathon {
	h := Hackathon{
		ID:                proto.ID,
		Name:              proto.Name,
		Description:       proto.Description,
		Location:          proto.Location,
		TeamSize:          proto.TeamSize,
		NumberOfTeams:     proto.NumberOfTeams,
		SubscriptionsOpen: proto.SubscriptionsOpen,
	}

	startDate, err := time.Parse("02/01/2006", proto.StartDate)
	if err == nil {
		h.StartDate = startDate
	}

	return &h
}

// Proto creates a protobuf representation of a hacakthon
func (h *Hackathon) Proto() *pb.Hackathon {
	return &pb.Hackathon{
		ID:                h.ID,
		Name:              h.Name,
		Description:       h.Description,
		Location:          h.Location,
		StartDate:         h.StartDate.Format("02/01/2006"), // TODO timestamp
		TeamSize:          h.TeamSize,
		NumberOfTeams:     h.NumberOfTeams,
		SubscriptionsOpen: h.SubscriptionsOpen,
	}
}

// Validate the hackathon's fields correcteness
func (h *Hackathon) Validate() error {
	if strings.TrimSpace(h.Name) == "" {
		return errors.New("name should not be empty")
	}
	if strings.TrimSpace(h.Location) == "" {
		return errors.New("location should not be empty")
	}
	if h.StartDate.IsZero() {
		return errors.New("invalid start date")
	}
	if h.TeamSize < 1 {
		return errors.New("team size should be greater than zero")
	}
	if h.NumberOfTeams < 1 {
		return errors.New("number of teams should be greater than zero")
	}

	return nil
}

// MarshalJSON encode the hackathon as JSON
func (h *Hackathon) MarshalJSON() ([]byte, error) {
	type Alias Hackathon
	return json.Marshal(&struct {
		StartDate string `json:"startDate"`
		*Alias
	}{
		StartDate: h.StartDate.Format("02/01/2006"),
		Alias:     (*Alias)(h),
	})
}

// UnmarshalJSON populate hackathon from JSON data
func (h *Hackathon) UnmarshalJSON(data []byte) error {
	type Alias Hackathon
	aux := &struct {
		StartDate string `json:"startDate"`
		*Alias
	}{
		Alias: (*Alias)(h),
	}

	var err error

	err = json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	h.StartDate, err = time.Parse("02/01/2006", aux.StartDate)
	if err != nil {
		return err
	}

	return nil
}
