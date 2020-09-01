package main

import (
	"fmt"
	"log"
	"net"

	"github.com/go-pg/pg"
	"github.com/thiagotbl/hackatons-api/api"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
	pb "github.com/thiagotbl/hackatons-api/protos"
	"github.com/thiagotbl/hackatons-api/repositories"
)

// TODO router config file
// TODO isolate dependency injection
func main() {
	db := connectDB()
	defer db.Close()

	hackathonsRepo := repositories.NewHackathonsRepository(db)
	// controller := controllers.HackathonsController{Repository: &repository}

	// r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	hackathons, _ := repository.GetHackathons()

	// 	for _, hackathon := range hackathons {
	// 		fmt.Fprintln(w, hackathon.Name+", "+hackathon.Location)
	// 	}
	// })
	// r.Get("/hackathons", controller.AllHackathons)
	// r.Post("/hackathons", controller.CreateHackathon)

	// http.ListenAndServe(":8080", r)

	organizerServer := api.NewOrganizerServer(hackathonsRepo)
	participantServer := api.NewParticipantServer(hackathonsRepo)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer()

	pb.RegisterOrganizerServiceServer(gRPCServer, organizerServer)
	pb.RegisterParticipantServiceServer(gRPCServer, participantServer)

	gRPCServer.Serve(listener)
}

func connectDB() *pg.DB {
	opt, err := pg.ParseURL("postgresql://postgres:passwd@localhost/hackathons_api?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	return db
}
