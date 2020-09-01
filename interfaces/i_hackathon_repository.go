package interfaces

import (
	"github.com/thiagotbl/hackatons-api/models"
)

// IHackathonRepository defines an interface for a hackathons repository
type IHackathonRepository interface {
	GetHackathons() ([]models.Hackathon, error)
	GetHackathon(id int64) (models.Hackathon, error)
	SaveHackathon(*models.Hackathon) error
	UpdateHackathon(*models.Hackathon) error
	DeleteHackathon(id int64) error
	GetHackathonTeams(hackathonID int64) ([]models.Team, error)
	GetTeamsByNameAsc(hackathonID int64, page int, limit int) ([]models.Team, error)
	GetTeamsByNameDesc(hackathonID int64, page int, limit int) ([]models.Team, error)
	UpdateSubscriptionsState(id int64, subscriptionsOpen bool) error

	SaveTeam(*models.Team) error
	UpdateTeam(*models.Team) error
	GetTeam(id int64) (models.Team, error)
}
