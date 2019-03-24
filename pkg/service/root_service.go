package service

import (
	"github.com/jinzhu/gorm"

	"github.com/goforbroke1006/sport-archive-svc/pkg/domain"
)

type SportArchiveService interface {
	GetSport(name string) (*domain.Sport, error)
	GetParticipant(name string) (*domain.Participant, error)
}

type sportArchiveService struct {
	db            *gorm.DB
	allowSaveData bool
}

func (svc sportArchiveService) GetSport(name string) (*domain.Sport, error) {
	var sport domain.Sport
	svc.db.Where(&domain.Sport{Name: name}).First(&sport)
	return &sport, nil
}

func (svc sportArchiveService) GetParticipant(name string) (*domain.Participant, error) {
	panic("implement me")
}

func NewSportArchiveService(db *gorm.DB, allowSave bool) SportArchiveService {
	return &sportArchiveService{
		db:            db,
		allowSaveData: allowSave,
	}
}
