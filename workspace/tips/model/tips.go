package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"io/ioutil"
	"strings"
)

type Tips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

//GetTip returning Tip/Command According to each title
func GetTip(title string) string {
	//data, err := loadTipsFromJson("../data/tips.json")
	data, err := loadTipsFromJson()
	if err != nil {
		return err.Error()
	} else if title != "" {
		commands := getAllCommands(data, title)
		for _, tip := range commands {
			return tip
		}
	} else if title == "" {
		return "should not be Empty"
	}
	return "Tips Not Available for Topic"
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

//reading json data from file
func readJsonFile(path string) ([]byte, error) {
	var errFileNotFound = errors.New("file not found")
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errFileNotFound
	}
	return jsonData, nil
}

//loading json data into Tips struct
func loadTipsFromJson() ([]Tips, error) {
	// run an app from main.go -> file path should be "data/tips.json"
	// if want to check all unit test cases ->file path should be "../data/tips.json"
	var path = getJsonFilePath()
	fmt.Println("---->" + path)
	var errorFile = errors.New("failed loading jSON file")
	data, err := readJsonFile(path)
	if err != nil {
		return nil, errorFile
	}
	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result, nil
}

func getJsonFilePath() string {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// remove base directory from the workingDir when run from test
	baseDir := filepath.Base(workingDir)
	isInTest := os.Getenv("GO_ENV") == "test"
	if isInTest {
		workingDir = strings.ReplaceAll(workingDir, baseDir, "")
	}
	return workingDir + "/data/tips.json"
}
