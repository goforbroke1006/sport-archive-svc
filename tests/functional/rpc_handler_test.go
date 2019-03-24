package functional

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"

	"github.com/goforbroke1006/sport-archive-svc/pkg/domain"
	"github.com/goforbroke1006/sport-archive-svc/pkg/handler"
	"github.com/goforbroke1006/sport-archive-svc/pkg/service"
)

const handlerAddr = "localhost:1234"

func TestMain(m *testing.M) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&domain.Sport{})
	db.AutoMigrate(&domain.Participant{})

	svc := service.NewSportArchiveService(db, true)

	handler.HandleClientsRequests(handlerAddr, svc)
}

func TestGetSportMethod(t *testing.T) {
	data := []byte(`{"name": "wildfowl"}`)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://"+handlerAddr, "application/json", r)
	if err != nil {
		t.Error(err)
	}
	resultData, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		t.Error(err)
	}

	if id, err := strconv.Atoi(string(resultData)); id <= 0 || err != nil {
		t.Error("it does not work")
	}
}
