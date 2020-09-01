package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func validHackathon() Hackathon {
	return Hackathon{
		ID:                1,
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

	require.Nil(t, err)
}

func TestValidateNoName(t *testing.T) {
	hackathon := validHackathon()
	hackathon.Name = ""

	err := hackathon.Validate()

	require.NotNil(t, err, "hackathon with no name should return an Error")
	require.Equal(t, "name should not be empty", err.Error())
}

func TestValidateNoLocation(t *testing.T) {
	hackathon := validHackathon()
	hackathon.Location = ""

	err := hackathon.Validate()

	require.NotNil(t, err, "hackathon with no location should return an error")
	require.Equal(t, "location should not be empty", err.Error())
}

func TestValidateNoStartDate(t *testing.T) {
	hackathon := validHackathon()
	hackathon.StartDate = time.Time{}

	err := hackathon.Validate()

	require.NotNil(t, err, "hackathon with an invalid start date should return an error")
	require.Equal(t, "invalid start date", err.Error())
}

func TestValidateInvalidTeamSize(t *testing.T) {
	hackathon := validHackathon()
	hackathon.TeamSize = 0

	err := hackathon.Validate()

	require.NotNil(t, err, "hackathon with invalid team size should return an error")
	require.Equal(t, "team size should be greater than zero", err.Error())
}

func TestValidateInvalidNumberOfTeams(t *testing.T) {
	hackathon := validHackathon()
	hackathon.NumberOfTeams = 0

	err := hackathon.Validate()

	require.NotNil(t, err, "hackathon with invalid number of teams should return an error")
	require.Equal(t, "number of teams should be greater than zero", err.Error())
}
