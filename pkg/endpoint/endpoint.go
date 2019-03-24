package endpoint

import (
	"errors"
	"net/http"

	"github.com/goforbroke1006/sport-archive-svc/pkg/service"
)

type SportArchiveServiceEndpoint struct {
	svc service.SportArchiveService
}

func (ep SportArchiveServiceEndpoint) GetSport(
	req *http.Request, args *GetSportArgs, res *GetSportResult,
) error {
	sport, err := ep.svc.GetSport(args.Name)
	if nil != err {
		return err
	}

	if sport == nil {
		return errors.New("not found")
	}

	res.ID = sport.ID
	res.Name = sport.Name

	return nil
}

func (ep SportArchiveServiceEndpoint) GetParticipant(
	req *http.Request, args *GetParticipantArgs, res *GetParticipantResult,
) error {

	prt, err := ep.svc.GetParticipant(args.Name, args.SportName)
	if nil != err {
		return err
	}

	res.ID = prt.ID
	res.Name = prt.Name
	res.Type = prt.Type
	res.SportID = prt.SportID

	return nil
}

func NewSportArchiveService(svc service.SportArchiveService) *SportArchiveServiceEndpoint {
	return &SportArchiveServiceEndpoint{
		svc: svc,
	}
}
