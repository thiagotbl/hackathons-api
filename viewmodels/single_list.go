package viewmodels

import (
	"github.com/thiagotbl/hackatons-api/models"
)

type SingleList struct {
	Data []models.Hackathon `json:"data"`
}
