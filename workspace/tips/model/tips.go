package model

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// tips class with field(title and tip)
type Tips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

const (
	default_value     = "invalid command ,please pass valid tool command "
	empty_stringVAlue = " "
)

//GetTip returning Tip/Command to the controller
func GetTip(title string) string {
	data, _ := loadTipsFromJson()
	commands := getAllCommands(data, title)
	for _, tip := range commands {
		return tip
	}
	return default_value
}

type dataInfo struct {
	titleData string
	tipData   string
}

//getting all tips and titles
func getAllCommands(data []Tips, title string) []string {
	title = title + empty_stringVAlue
	cmd := strings.Split(title, empty_stringVAlue)
	cmds := []dataInfo{}
	subcommands := make([]string, 0)

	for _, value := range data {
		details := dataInfo{}
		if strings.Contains(value.Tip, cmd[0]) {
			details = dataInfo{titleData: value.Title, tipData: value.Tip}
			cmds = append(cmds, details)
		}
	}
	for _, val := range cmds {
		if strings.Contains(val.tipData, cmd[1]) {
			command := val.titleData + " : " + val.tipData
			subcommands = append(subcommands, command)
		}
	}
	return subcommands
}

//unmarshal json data into Tips struct
func loadTipsFromJson() ([]Tips, error) {
	// run an app from main.go -> file path should be "data/tips.json"
	// if want to check all unit test cases ->file path should be "../data/tips.json"
	var path = getJsonFilePath()
	var data []byte
	data, _ = readJsonFile(path)
	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result, nil
}

// getting file path for main file and testing function
func getJsonFilePath() string {
	currentDir, _ := getCurrentWorkingDir()
	// remove base directory from the workingDir when run from test
	baseDir := filepath.Base(currentDir)

	isInTest := os.Getenv("GO_ENV") == "test"
	if isInTest {
		currentDir = strings.ReplaceAll(currentDir, baseDir, "")

	}
	return currentDir + "/data/tips.json"
}

// get json file data
var fileRead = os.ReadFile

//reading data from json file
func readJsonFile(path string) ([]byte, error) {
	data, err := fileRead(path)
	if err != nil {
		logrus.WithField("file path ", path).Debug("unsuccessfully reading the file path ")
		return nil, err
	}
	logrus.WithField("file path ", path).Debug("successfully reading the file path ")
	return data, nil
}

// Get Working directory function
var osGetWd = os.Getwd

//getting current working dir.
func getCurrentWorkingDir() (string, error) {
	workingDir, err := osGetWd()
	if err != nil {
		logrus.WithField("working dir", workingDir).Debug("unsuccessfully reading the working dir path ")
		return "", errors.New("could not get current working directory")
	}
	logrus.WithField("working dir", workingDir).Debug("successfully reading the working dir path ")
	return workingDir, nil
}
