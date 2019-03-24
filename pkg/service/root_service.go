package service

import (
	"github.com/google/logger"
	"github.com/jinzhu/gorm"
	"strings"

	"github.com/goforbroke1006/sport-archive-svc/pkg/dao"
	"github.com/goforbroke1006/sport-archive-svc/pkg/domain"
)

type SportArchiveService interface {
	GetSport(name string) (*domain.Sport, error)
	GetParticipant(name, sportName string) (*domain.Participant, error)
}

type sportArchiveService struct {
	db            *gorm.DB
	allowSaveData bool
}

func (svc sportArchiveService) GetSport(name string) (*domain.Sport, error) {
	logger.Infof("Looking for sport: %s", name)

	name = strings.ToLower(name)

	sport, err := dao.GetSportByName(svc.db, name)
	if nil != err && err.Error() != "record not found" {
		return nil, err
	}
	if nil == sport && svc.allowSaveData {
		logger.Infof("Create sport: %s", name)
		sport = &domain.Sport{Name: name}
		if err := dao.CreateSport(svc.db, sport); nil != err {
			return nil, err
		}
	}
	if result := svc.db.Where(&domain.Sport{Name: name}).First(&sport); result.Error != nil {
		return nil, result.Error
	}
	return sport, nil
}

func (svc sportArchiveService) GetParticipant(name, sportName string) (*domain.Participant, error) {
	sportName = strings.ToLower(sportName)
	sport, err := svc.GetSport(sportName)
	if nil != err {
		return nil, err
	}

	logger.Infof("Looking for participant: %s", name)

	name = strings.ToLower(name)

	participant, err := dao.GetParticipantByName(svc.db, name, sport.ID)
	if nil != err && err.Error() != "record not found" {
		return nil, err
	}
	if nil == participant && svc.allowSaveData {
		logger.Infof("Create participant: %s", name)
		participant = &domain.Participant{
			Name:    name,
			SportID: sport.ID,
			Sport:   *sport,
		}
		if err := dao.CreateParticipant(svc.db, participant); nil != err {
			return nil, err
		}
	}

	return participant, nil
}

func NewSportArchiveService(db *gorm.DB, allowSave bool) SportArchiveService {
	return &sportArchiveService{
		db:            db,
		allowSaveData: allowSave,
	}
}
