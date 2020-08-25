package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thiagotbl/hackatons-api/interfaces"
	"github.com/thiagotbl/hackatons-api/viewmodels"
)

type HackathonsController struct {
	interfaces.IHackathonRepository
}

// TODO golint to limit line size?
func (controller *HackathonsController) AllHackathons(w http.ResponseWriter, req *http.Request) {
	// TODO handle errors
	hackathons, _ := controller.GetHackathons()
	view := viewmodels.SingleList{Data: hackathons}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(view)
}
