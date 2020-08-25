package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/thiagotbl/hackatons-api/controllers"

	"github.com/thiagotbl/hackatons-api/repositories"

	"github.com/thiagotbl/hackatons-api/infra"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

// TODO router config file
// TODO isolate dependency injection
func main() {
	db := connectDB()
	defer db.Close()

	postgresDB := infra.PostgresDb{Conn: db}
	repository := repositories.HackathonRepository{IDb: &postgresDB}
	controller := controllers.HackathonsController{&repository}

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		hackathons, _ := repository.GetHackathons()

		for _, hackathon := range hackathons {
			fmt.Fprintln(w, hackathon.Name+", "+hackathon.Location)
		}
	})
	r.Get("/hackathons", controller.AllHackathons)

	http.ListenAndServe(":8080", r)
}

func connectDB() *sql.DB {
	db, err := sql.Open(
		"postgres",
		"postgresql://postgres:passwd@localhost/hackathons_api?sslmode=disable",
	)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("db ping failed: ", err)
	}

	return db
}
