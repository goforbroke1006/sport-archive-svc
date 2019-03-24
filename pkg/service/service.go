package service

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type GetSportArgs struct {
	Name string `json:"name"`
}

type GetSportResult struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

//type SportArchiveService interface {
//	GetSportID(c context.Context, request *GetSportArgs) error
//	GetTeamID(c context.Context, name string) uint64
//}

type SportArchiveService struct {
	db            *gorm.DB
	allowSaveData bool
}

func (svc SportArchiveService) GetSport(req *http.Request, args *GetSportArgs, res *GetSportResult) error {
	//log.Println(request.Name)
	res.ID = 666

	return nil
}

//func (svc SportArchiveService) GetTeamID(c context.Context, name string) uint64 {
//	return 0
//}

func NewSportArchiveService(db *gorm.DB, allowSave bool) *SportArchiveService {
	return &SportArchiveService{
		db:            db,
		allowSaveData: allowSave,
	}
}
