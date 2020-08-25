package repositories

import (
	"github.com/thiagotbl/hackatons-api/interfaces"
	"github.com/thiagotbl/hackatons-api/models"
)

type HackathonRepository struct {
	interfaces.IDb
}

// TODO maybe use https://github.com/jmoiron/sqlx to ease field mapping
func (repository *HackathonRepository) GetHackathons() ([]models.Hackathon, error) {
	rows, err := repository.Query("SELECT id, name, location FROM hackathons")
	if err != nil {
		return nil, err
	}

	var hackathons []models.Hackathon
	for rows.Next() {
		var hackathon models.Hackathon
		rows.Scan(&hackathon.Id, &hackathon.Name, &hackathon.Location)

		hackathons = append(hackathons, hackathon)
	}

	return hackathons, nil
}
