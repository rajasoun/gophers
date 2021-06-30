package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"strings"
)

type Tips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

const (
	empty_string  = ""
	empty_value   = "should not be Empty"
	default_value = "Tips Not Available for Topic"
)

//GetTip returning Tip/Command According to each title
func GetTip(title string, model Model) string {
	data, _ := loadTipsFromJson(model)
	if title != empty_string {
		commands := getAllCommands(data, title)
		for _, tip := range commands {
			return tip
		}
	} else if title == empty_string {
		return empty_value
	}
	return default_value
}

func getAllCommands(data []Tips, title string) []string {
	commands := make([]string, 0)
	for index := range data {
		if strings.Contains(data[index].Tip, title) {
			command := data[index].Title + " : " + data[index].Tip
			commands = append(commands, command)
		}
	}
	return commands
}

//loading json data into Tips struct
func loadTipsFromJson(model Model) ([]Tips, error) {
	// run an app from main.go -> file path should be "data/tips.json"
	// if want to check all unit test cases ->file path should be "../data/tips.json"
	var path = getJsonFilePath()
	var data []byte
	data, _ = readJsonFile(path, model)
	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result, nil
}
func getJsonFilePath() string {
	currentDir, _ := getCurrentWorkingDir(&handlerImpl{})
	// remove base directory from the workingDir when run from test
	baseDir := filepath.Base(currentDir)
	isInTest := os.Getenv("GO_ENV") == "test"
	if isInTest {
		currentDir = strings.ReplaceAll(currentDir, baseDir, "")
	}
	return currentDir + "/data/tips.json"
}

//readFile impl
type Model interface {
	readFile(path string) ([]byte, error)
}
type Reader struct{}

func (reader Reader) readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func readJsonFile(path string, model Model) ([]byte, error) {
	var errFileNotFound = errors.New("failed loading jSON file")
	jsonData, err := model.readFile(path)
	if err != nil {
		fmt.Println(errFileNotFound)
		return nil, errFileNotFound
	}
	return jsonData, nil
}

//getWd impl
type handler interface {
	get_wd() (string, error)
}
type handlerImpl struct{}

func (han_impl *handlerImpl) get_wd() (string, error) {
	return os.Getwd()
}
func getCurrentWorkingDir(handler handler) (string, error) {
	var workingDir string
	workingDir, _ = handler.get_wd()
	return workingDir, nil

}
