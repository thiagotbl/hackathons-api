package models

import (
	pb "github.com/thiagotbl/hackatons-api/protos"
)

// Team represents a hackathon team
type Team struct {
	ID          int64
	Name        string
	Confirmed   bool `sql:",notnull"`
	HackathonID int64
	// SubscriptionDate time.Time // TODO
	Participants []*Participant
}

// TeamFromProto transforms a *protos.Team into a Team struct
func TeamFromProto(proto *pb.Team) *Team {
	var members []*Participant
	for _, member := range proto.Participants {
		members = append(members, ParticipantFromProto(member))
	}

	team := Team{
		ID:           proto.ID,
		Name:         proto.Name,
		Confirmed:    proto.Confirmed,
		HackathonID:  proto.HackathonId,
		Participants: members,
	}

	return &team
}

// Proto creates a protobuf representation of a team
func (t *Team) Proto() *pb.Team {
	var members []*pb.Participant
	for _, member := range t.Participants {
		members = append(members, member.Proto())
	}

	return &pb.Team{
		ID:           t.ID,
		Name:         t.Name,
		Confirmed:    t.Confirmed,
		HackathonId:  t.HackathonID,
		Participants: members,
	}
}
