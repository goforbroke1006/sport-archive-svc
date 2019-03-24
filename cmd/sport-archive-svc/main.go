package main

import (
	"flag"
	"io"
	"os"

	"github.com/google/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"

	"github.com/goforbroke1006/sport-archive-svc/pkg/domain"
	"github.com/goforbroke1006/sport-archive-svc/pkg/endpoint"
	"github.com/goforbroke1006/sport-archive-svc/pkg/handler"
	"github.com/goforbroke1006/sport-archive-svc/pkg/service"
)

var (
	dbConnStr     = flag.String("db-conn", "./sport-archive.db", "")
	handleAddr    = flag.String("handle-addr", "127.0.0.1:10001", "")
	allowSaveData = flag.Bool("allow-save", false, "")
	logPath       = flag.String("log-path", "./access.log", "")
	verbose       = flag.Bool("verbose", true, "print info level logs to stdout")
)

func init() {
	flag.Parse()
}

func main() {
	logFile, err := os.OpenFile(*logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer finalizeCloser(logFile)

	defer logger.Init("sport-archive-svc", *verbose, true, logFile).Close()

	db, err := gorm.Open("sqlite3", *dbConnStr)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer finalizeCloser(db)
	db.AutoMigrate(&domain.Sport{})
	db.AutoMigrate(&domain.Participant{})

	svc := service.NewSportArchiveService(db, *allowSaveData)
	eps := endpoint.NewSportArchiveService(svc)
	handler.HandleClientsRequests(*handleAddr, eps)
}

func finalizeCloser(c io.Closer) {
	err := c.Close()
	if nil != err {
		logger.Fatalf("Failed to close descriptor: %v", err)
	}
}
