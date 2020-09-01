package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thiagotbl/hackatons-api/interfaces"
	"github.com/thiagotbl/hackatons-api/models"
	"github.com/thiagotbl/hackatons-api/views"
)

// TODO use a service instead the hackathons repository?
type HackathonsController struct {
	Repository interfaces.IHackathonRepository
}

func (c *HackathonsController) AllHackathons(w http.ResponseWriter, req *http.Request) {
	hackathons, err := c.Repository.GetHackathons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view := views.SingleList{Data: hackathons}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view)
}

func (c *HackathonsController) CreateHackathon(w http.ResponseWriter, req *http.Request) {
	var hackathon models.Hackathon
	err := json.NewDecoder(req.Body).Decode(&hackathon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = hackathon.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = c.Repository.SaveHackathon(&hackathon)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hackathon)
}
