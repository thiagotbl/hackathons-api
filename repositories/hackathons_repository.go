package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiagotbl/hackatons-api/interfaces"
	"github.com/thiagotbl/hackatons-api/models"
)

type hackathonsRepository struct {
	DB *pg.DB
}

// NewHackathonsRepository creates a new hackathons repository backed by go-pg
func NewHackathonsRepository(db *pg.DB) interfaces.IHackathonRepository {
	return &hackathonsRepository{DB: db}
}

// TODO order
// TODO pagination
func (r hackathonsRepository) GetHackathons() ([]models.Hackathon, error) {
	var hackathons []models.Hackathon

	err := r.DB.Model(&hackathons).Select()

	return hackathons, err
}

func (r hackathonsRepository) GetHackathon(id int64) (models.Hackathon, error) {
	hackathon := &models.Hackathon{ID: id}
	err := r.DB.Model(hackathon).WherePK().Select()

	return *hackathon, err
}

func (r hackathonsRepository) SaveHackathon(hackathon *models.Hackathon) error {
	_, err := r.DB.Model(hackathon).Insert()

	return err
}

func (r hackathonsRepository) UpdateHackathon(hackathon *models.Hackathon) error {
	_, err := r.DB.Model(hackathon).WherePK().Update()

	return err
}

func (r hackathonsRepository) DeleteHackathon(id int64) error {
	hackathon := &models.Hackathon{ID: id}
	_, err := r.DB.Model(hackathon).WherePK().Delete()

	return err
}

// TODO order by name - ascending and descending
// TODO pagination
// Go does not have default parameters...
// Go does not have strict enums...
func (r hackathonsRepository) GetHackathonTeams(hackathonID int64) ([]models.Team, error) {
	var teams []models.Team
	err := r.DB.Model(&teams).
		Where("hackathon_id = ?", hackathonID).
		Relation("Participants").
		Select()

	return teams, err
}

func (r hackathonsRepository) GetTeamsByNameAsc(hackathonID int64, page int, limit int) ([]models.Team, error) {
	var teams []models.Team
	err := r.DB.Model(&teams).
		Where("hackathon_id = ?", hackathonID).
		Order("name ASC").
		Limit(limit).
		Offset(page*limit - limit).
		Relation("Participants").
		Select()

	return teams, err
}

func (r hackathonsRepository) GetTeamsByNameDesc(hackathonID int64, page int, limit int) ([]models.Team, error) {
	var teams []models.Team
	err := r.DB.Model(&teams).
		Where("hackathon_id = ?", hackathonID).
		Order("name DESC").
		Limit(limit).
		Offset(page*limit - limit).
		Relation("Participants").
		Select()

	return teams, err
}

// func (r hackathonsRepository) getTeamsByName() (

// )

func (r hackathonsRepository) UpdateSubscriptionsState(
	id int64, subscriptionsOpen bool,
) error {
	hackathon := models.Hackathon{
		ID:                id,
		SubscriptionsOpen: subscriptionsOpen,
	}

	_, err := r.DB.Model(&hackathon).
		Set("subscriptions_open = ?subscriptions_open").
		Where("id = ?id").
		Update()

	return err
}

func (r *hackathonsRepository) GetTeam(id int64) (models.Team, error) {
	team := &models.Team{ID: id}
	err := r.DB.Model(team).WherePK().Select()

	return *team, err
}

func (r *hackathonsRepository) SaveTeam(team *models.Team) error {
	// TODO transaction
	_, err := r.DB.Model(team).Insert()
	if err != nil {
		return err
	}

	for _, participant := range team.Participants {
		participant.TeamID = team.ID
	}

	_, err = r.DB.Model(&team.Participants).Insert()

	return err
}

func (r *hackathonsRepository) UpdateTeam(team *models.Team) error {
	_, err := r.DB.Model(team).WherePK().Update()

	return err
}
