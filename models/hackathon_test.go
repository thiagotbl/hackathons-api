package models

import (
	"testing"
	"time"
)

func validHackathon() Hackathon {
	return Hackathon{
		Id:                1,
		Name:              "Sampa Hack",
		Description:       "Not just another hackathon!",
		Location:          "São Paulo",
		StartDate:         time.Now(),
		TeamSize:          5,
		NumberOfTeams:     20,
		SubscriptionsOpen: true,
	}
}

func TestValidateAllFieldsFilled(t *testing.T) {
	hackathon := validHackathon()
	err := hackathon.Validate()

	if err != nil {
		t.Errorf("a hackathon with all fields filled should be valid")
	}
}

func TestValidateNoName(t *testing.T) {
	hackathon := validHackathon()
	hackathon.Name = ""

	err := hackathon.Validate()

	if err == nil {
		t.Errorf("hackaton with no name should return an error")
	}
	if err != nil && err.Error() != "name should not be empty" {
		t.Errorf(
			`unexpected error message: got "%v" want "%v"`,
			err.Error(),
			"name should not be empty",
		)
	}
}

func TestValidateNoLocation(t *testing.T) {
	hackathon := validHackathon()
	hackathon.Location = ""

	err := hackathon.Validate()

	if err == nil {
		t.Errorf("hackathon with no location should return an error")
	}
	if err != nil && err.Error() != "location should not be empty" {
		t.Errorf(
			`unexpected error: got "%v" want "%v"`,
			err.Error(),
			"location should not be empty",
		)
	}
}

func TestValidateNoStartDate(t *testing.T) {
	hackathon := validHackathon()
	hackathon.StartDate = time.Time{}

	err := hackathon.Validate()

	if err == nil {
		t.Errorf("hackathon with an invalid start date should return an error")
	}
	if err != nil && err.Error() != "invalid start date" {
		t.Errorf(
			`unexpected error: got "%v" want "%v"`,
			err.Error(),
			"invalid start date",
		)
	}
}

func TestValidateInvalidTeamSize(t *testing.T) {
	hackathon := validHackathon()
	hackathon.TeamSize = 0

	err := hackathon.Validate()

	if err == nil {
		t.Errorf("hackathon with invalid team size should return an error")
	}
	if err != nil && err.Error() != "team size should be greater than zero" {
		t.Errorf(
			`unexpected error: got "%v" want "%v"`,
			err.Error(),
			"team size should be greater than zero",
		)
	}
}

func TestValidateInvalidNumberOfTeams(t *testing.T) {
	hackathon := validHackathon()
	hackathon.NumberOfTeams = 0

	err := hackathon.Validate()

	if err == nil {
		t.Errorf("hackathon with invalid number of teams should return an error")
	}
	if err != nil && err.Error() != "number of teams should be greater than zero" {
		t.Errorf(
			`unexpected error: got "%v" want "%v"`,
			err.Error(),
			"number of teams should be greater than zero",
		)
	}
}
