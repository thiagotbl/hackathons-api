package api

import (
	"context"

	"github.com/thiagotbl/hackatons-api/interfaces"
	"github.com/thiagotbl/hackatons-api/models"
	pb "github.com/thiagotbl/hackatons-api/protos"
)

type organizerServer struct {
	Repository interfaces.IHackathonRepository
}

// NewOrganizerServer returns an organizer server
func NewOrganizerServer(r interfaces.IHackathonRepository) pb.OrganizerServiceServer {
	return &organizerServer{Repository: r}
}

func (s *organizerServer) CreateHackathon(
	ctx context.Context,
	req *pb.CreateHackathonRequest,
) (*pb.CreateHackathonResponse, error) {
	hackathon := models.HackathonFromProto(req.Hackathon)
	err := s.Repository.SaveHackathon(hackathon)

	return &pb.CreateHackathonResponse{Hackathon: hackathon.Proto()}, err
}

func (s *organizerServer) GetHackathons(
	context.Context, *pb.GetHackathonsRequest,
) (*pb.GetHackathonsResponse, error) {
	hackathons, err := s.Repository.GetHackathons()
	if err != nil {
		return nil, err
	}

	var response pb.GetHackathonsResponse
	for _, hackathon := range hackathons {
		response.Hackathons = append(response.Hackathons, hackathon.Proto())
	}

	return &response, nil
}

func (s *organizerServer) UpdateHackathon(
	ctx context.Context,
	req *pb.UpdateHackathonRequest,
) (*pb.UpdateHackathonResponse, error) {
	hackathon := models.HackathonFromProto(req.Hackathon)
	err := s.Repository.UpdateHackathon(hackathon)

	return &pb.UpdateHackathonResponse{Hackathon: hackathon.Proto()}, err
}

func (s *organizerServer) CloseSubscriptions(
	ctx context.Context, msg *pb.CloseSubscriptionsRequest,
) (*pb.CloseSubscriptionsResponse, error) {
	err := s.Repository.UpdateSubscriptionsState(msg.ID, false)

	return &pb.CloseSubscriptionsResponse{}, err
}

func (s *organizerServer) OpenSubscriptions(
	ctx context.Context, msg *pb.OpenSubscriptionsRequest,
) (*pb.OpenSubscriptionsResponse, error) {
	err := s.Repository.UpdateSubscriptionsState(msg.ID, true)

	return &pb.OpenSubscriptionsResponse{}, err
}

func (s *organizerServer) DeleteHackathon(
	ctx context.Context, msg *pb.DeleteHackathonRequest,
) (*pb.DeleteHackathonResponse, error) {
	err := s.Repository.DeleteHackathon(msg.ID)

	return &pb.DeleteHackathonResponse{}, err
}

func (s *organizerServer) GetHackathonTeams(
	ctx context.Context, msg *pb.GetHackathonTeamsRequest,
) (*pb.GetHackathonTeamsResponse, error) {
	teams, err := s.Repository.GetHackathonTeams(msg.HackathonId)
	if err != nil {
		return nil, err
	}

	var response pb.GetHackathonTeamsResponse
	for _, team := range teams {
		response.Teams = append(response.Teams, team.Proto())
	}

	return &response, nil
}

func (s *organizerServer) GetHackathonTeamsByName(
	ctx context.Context, msg *pb.GetHackathonTeamsByNameRequest,
) (*pb.GetHackathonTeamsByNameResponse, error) {
	hackathonID := msg.GetHackathonId()
	order := msg.GetOrder().String()
	page := int(msg.GetPage())
	limit := int(msg.GetLimit())

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	var teams []models.Team
	var err error
	if order == "ASC" {
		teams, err = s.Repository.GetTeamsByNameAsc(hackathonID, page, limit)
	} else {
		teams, err = s.Repository.GetTeamsByNameDesc(hackathonID, page, limit)
	}

	if err != nil {
		return nil, err
	}

	var response pb.GetHackathonTeamsByNameResponse
	for _, team := range teams {
		response.Teams = append(response.Teams, team.Proto())
	}

	return &response, nil
}
