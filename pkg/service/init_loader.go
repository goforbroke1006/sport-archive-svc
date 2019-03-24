package service

import (
	"bufio"
	"bytes"
	"github.com/goforbroke1006/sport-archive-svc/pkg/dao"
	"github.com/goforbroke1006/sport-archive-svc/pkg/domain"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
)

type InitLoader interface {
	InitSportsList(filename string) error
	InitParticipantsList(filename string) error
}

type initLoader struct {
	db *gorm.DB
}

func (ldr initLoader) InitSportsList(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if nil != err {
		return err
	}
	buffer := bytes.NewBuffer(data)
	reader := bufio.NewReader(buffer)

	for {
		line, _, err := reader.ReadLine()
		if nil != err {
			break
		}

		lineStr := string(line)

		err = dao.CreateSport(ldr.db, &domain.Sport{Name: lineStr})
		if err != nil {
			log.Println("Error:", err.Error())
		}
	}

	return nil
}

func (ldr initLoader) InitParticipantsList(filename string) error {
	return nil
}

func NewInitLoader(db *gorm.DB) InitLoader {
	return &initLoader{
		db: db,
	}
}
