package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/google/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"

	"github.com/goforbroke1006/sport-archive-svc/pkg/domain"
	"github.com/goforbroke1006/sport-archive-svc/pkg/endpoint"
	"github.com/goforbroke1006/sport-archive-svc/pkg/handler"
	"github.com/goforbroke1006/sport-archive-svc/pkg/http"
	"github.com/goforbroke1006/sport-archive-svc/pkg/service"
	"github.com/goforbroke1006/sport-archive-svc/pkg/trace"
)

var (
	dbConnStr     = flag.String("db-conn", "./sport-archive.db", "")
	serveAddr     = flag.String("serve-addr", "0.0.0.0:8080", "")
	allowSaveData = flag.Bool("allow-save", true, "")
	logPath       = flag.String("log-path", "./access.log", "")
	verbose       = flag.Bool("verbose", true, "Print info level logs to stdout")

	zipkinAddr = flag.String("zipkin-addr", "http://127.0.0.1:9411", "Select zipkin host")
)

const serviceName = "sport-archive-svc"

func init() {
	flag.Parse()
}

func main() {
	logFile, err := os.OpenFile(*logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer finalizeCloser(logFile)

	defer logger.Init(serviceName, *verbose, false, logFile).Close()

	port := http.ParsePortFromAddr(*serveAddr)
	tracer, err := trace.NewTracer(*zipkinAddr, serviceName, port)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("sqlite3", *dbConnStr)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer finalizeCloser(db)
	db.AutoMigrate(&domain.Sport{})
	db.AutoMigrate(&domain.Participant{})

	svc := service.NewSportArchiveService(db, *allowSaveData)
	eps := endpoint.NewSportArchiveService(svc)
	handler.HandleClientsRequests(
		*serveAddr,
		eps,
		zipkinhttp.NewServerMiddleware(
			tracer,
			zipkinhttp.SpanName("request"),
		),
	)
}

func finalizeCloser(c io.Closer) {
	err := c.Close()
	if nil != err {
		logger.Fatalf("Failed to close descriptor: %v", err)
	}
}
