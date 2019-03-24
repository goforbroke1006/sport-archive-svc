package service

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jinzhu/gorm"
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
		parts := strings.Split(lineStr, " ")

		if len(parts) < 2 {
			return errors.New(fmt.Sprintf("wrong args for sport initialization : %s", lineStr))
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
