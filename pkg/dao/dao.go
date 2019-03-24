package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/goforbroke1006/sport-archive-svc/pkg/domain"
)

func CreateSport(db *gorm.DB, sport *domain.Sport) error {
	result := db.Create(sport)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetSportByName(db *gorm.DB, name string) (*domain.Sport, error) {
	var sport domain.Sport
	result := db.Where(&domain.Sport{Name: name}).First(&sport)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sport, nil
}

func CreateParticipant(db *gorm.DB, ptn *domain.Participant) error {
	result := db.Create(domain.Participant{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetParticipantByName(db *gorm.DB, name string) (*domain.Participant, error) {
	var ptn domain.Participant
	result := db.Where(&domain.Participant{Name: name}).First(&ptn)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ptn, nil
}
