package models

import (
	pb "github.com/thiagotbl/hackatons-api/protos"
)

// T-shirt Sizes
const (
	Small      = "S"
	Medium     = "M"
	Large      = "L"
	ExtraLarge = "XL"
)

// Participant represents a team member
type Participant struct {
	ID         int64
	Name       string
	Email      string
	PhotoURL   string
	Phone      string
	TShirtSize string
	TeamID     int64
	Team       *Team
}

// ParticipantFromProto transforms a *protos.Participant into a Participant struct
func ParticipantFromProto(proto *pb.Participant) *Participant {
	participant := Participant{
		ID:         proto.ID,
		Name:       proto.Name,
		Email:      proto.Email,
		PhotoURL:   proto.PhotoUrl,
		Phone:      proto.Phone,
		TShirtSize: proto.TShirtSize.String(),
		TeamID:     proto.TeamId,
	}

	return &participant
}

// Proto creates a protobuf representation of a participant
func (p *Participant) Proto() *pb.Participant {
	return &pb.Participant{
		ID:         p.ID,
		Name:       p.Name,
		Email:      p.Email,
		PhotoUrl:   p.PhotoURL,
		Phone:      p.Phone,
		TShirtSize: pb.T_SHIRT_SIZE(pb.T_SHIRT_SIZE_value[p.TShirtSize]),
		TeamId:     p.TeamID,
	}
}
