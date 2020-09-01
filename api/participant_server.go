package api

import (
	"context"

	i "github.com/thiagotbl/hackatons-api/interfaces"
	"github.com/thiagotbl/hackatons-api/models"
	pb "github.com/thiagotbl/hackatons-api/protos"
)

type participantServer struct {
	HackathonsRepository i.IHackathonRepository
}

// NewParticipantServer returns an participant server
func NewParticipantServer(hackathonsRepo i.IHackathonRepository) pb.ParticipantServiceServer {
	return &participantServer{
		HackathonsRepository: hackathonsRepo,
	}
}

func (s *participantServer) ListHackathons(
	context.Context, *pb.ListHackathonsRequest,
) (*pb.ListHackathonsResponse, error) {
	hackathons, err := s.HackathonsRepository.GetHackathons()
	if err != nil {
		return nil, err
	}

	var response pb.ListHackathonsResponse
	for _, hackathon := range hackathons {
		response.Hackathons = append(response.Hackathons, hackathon.Proto())
	}

	return &response, nil
}

// TODO move logic from here - a service that should be injected on this "controller"
// TODO I'm not happy with this method. How to write this better?
func (s *participantServer) SubscribeTeam(
	ctx context.Context, msg *pb.SubscribeTeamRequest,
) (*pb.SubscribeTeamResponse, error) {
	hackathon, err := s.HackathonsRepository.GetHackathon(msg.GetTeam().GetHackathonId())
	if err != nil {
		return nil, err
	}

	if hackathon.SubscriptionsOpen == false {
		errorMessage := pb.Error{
			Code:    "422",
			Message: "Hackathon subscriptions are closed",
		}
		response := pb.SubscribeTeamResponse{Error: &errorMessage}

		return &response, nil
	}

	team := models.TeamFromProto(msg.GetTeam())

	if len(team.Participants) != int(hackathon.TeamSize) {
		errorMessage := pb.Error{
			Code:    "422",
			Message: "Invalid team size",
		}

		return &pb.SubscribeTeamResponse{Error: &errorMessage}, nil
	}

	err = s.HackathonsRepository.SaveTeam(team)

	return &pb.SubscribeTeamResponse{}, err
}

func (s *participantServer) CancelTeamSubscription(
	ctx context.Context, msg *pb.CancelTeamSubscriptionRequest,
) (*pb.CancelTeamSubscriptionResponse, error) {
	team, err := s.HackathonsRepository.GetTeam(msg.ID)
	if err != nil {
		return nil, err
	}

	team.Confirmed = false
	err = s.HackathonsRepository.UpdateTeam(&team)

	return &pb.CancelTeamSubscriptionResponse{}, err
}
