package domain

import "database/sql"

const (
	ParticipantTypeCountry = "country"
	ParticipantTypeTeam    = "team"
	ParticipantTypePlayer  = "player"
)

type Sport struct {
	ID   uint64 `gorm:"not null;unique"`
	Name string `gorm:"not null;unique"`
}

type Participant struct {
	ID      uint64        `gorm:"not null;unique"`
	SportID sql.NullInt64 `gorm:"index:participant_data"`
	Sport   Sport
	Name    string `gorm:"index:participant_data"`
	Type    string `gorm:"index:participant_data"`
}
