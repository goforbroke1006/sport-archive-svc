package main

import (
	"flag"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"

	"github.com/goforbroke1006/sport-archive-svc/pkg/domain"
	"github.com/goforbroke1006/sport-archive-svc/pkg/service"
)

var (
	dbConnStr               = flag.String("db-conn", "./sport-archive.db", "")
	sportsFixturePath       = flag.String("sport-fixture", "", "")
	participantsFixturePath = flag.String("participant-fixture", "", "")
)

func init() {
	flag.Parse()
}

func main() {
	db, err := gorm.Open("sqlite3", *dbConnStr)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&domain.Sport{})
	db.AutoMigrate(&domain.Participant{})

	loader := service.NewInitLoader(db)
	if nil != sportsFixturePath && len(*sportsFixturePath) > 0 {
		err = loader.InitSportsList(*sportsFixturePath)
		if nil != err {
			log.Fatalf("Failed load sports fixture '%s' %s", *sportsFixturePath, err)
		}
	}
	if nil != participantsFixturePath && len(*participantsFixturePath) > 0 {
		err = loader.InitParticipantsList(*participantsFixturePath)
		if nil != err {
			log.Fatalf("Failed load participants fixture '%s' %s", *participantsFixturePath, err)
		}
	}
}
