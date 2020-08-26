package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/google/go-cmp/cmp"
	"github.com/thiagotbl/hackatons-api/models"
	"github.com/thiagotbl/hackatons-api/viewmodels"
)

var fakeHackathons = []models.Hackathon{
	models.Hackathon{
		Id:          12,
		Name:        "Fake Hack!",
		Description: "Let's fake it!",
		Location:    "São Paulo",
	},
	models.Hackathon{
		Id:          13,
		Name:        "A Very Real Hackathon!",
		Description: "No fake at all!",
		Location:    "São Paulo",
	},
}

type fakeRepository struct{}

func (*fakeRepository) GetHackathons() ([]models.Hackathon, error) {
	return fakeHackathons, nil
}

type brokenRepository struct{}

func (*brokenRepository) GetHackathons() ([]models.Hackathon, error) {
	return nil, errors.New("something went wrong reading hackathons")
}

func allHackathonsRequest() *httptest.ResponseRecorder {
	subject := HackathonsController{&fakeRepository{}}

	req := httptest.NewRequest("GET", "/hackathons", nil)
	recorder := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/hackathons", subject.AllHackathons)
	r.ServeHTTP(recorder, req)

	return recorder
}

func TestAllHackathonsReturnsCorrectHeaders(t *testing.T) {
	recorder := allHackathonsRequest()

	if recorder.Code != http.StatusOK {
		t.Errorf(
			"wrong status code: got %v want %v",
			recorder.Code,
			http.StatusOK,
		)
	}
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Errorf(
			"wrong Content-Type: got %v want %v",
			recorder.Header().Get("Content-Type"),
			"application/json",
		)
	}
}

func TestAllHackathonsReturnsCorrectData(t *testing.T) {
	recorder := allHackathonsRequest()

	var result viewmodels.SingleList
	json.NewDecoder(recorder.Body).Decode(&result)

	expectedResult := viewmodels.SingleList{Data: fakeHackathons}

	if !cmp.Equal(result, expectedResult) {
		t.Errorf(
			"invalid hackathons list: got %v want %v",
			result,
			expectedResult,
		)
	}
}

func TestAllHackathonsRepositoryFails(t *testing.T) {
	subject := HackathonsController{&brokenRepository{}}

	req := httptest.NewRequest("GET", "/hackathons", nil)
	recorder := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/hackathons", subject.AllHackathons)
	r.ServeHTTP(recorder, req)

	bodyBytes, _ := ioutil.ReadAll(recorder.Body)
	message := strings.TrimSpace(string(bodyBytes))

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf(
			"wrong status code: got %v want %v",
			recorder.Code,
			http.StatusInternalServerError,
		)
	}

	if message != "something went wrong reading hackathons" {
		t.Errorf(
			`unexpected error message: got "%v" want "%v"`,
			message,
			"something went wrong reading hackathons",
		)
	}
}
