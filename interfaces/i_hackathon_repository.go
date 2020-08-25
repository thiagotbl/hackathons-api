package interfaces

import (
	"github.com/thiagotbl/hackatons-api/models"
)

type IHackathonRepository interface {
	GetHackathons() ([]models.Hackathon, error)
}
