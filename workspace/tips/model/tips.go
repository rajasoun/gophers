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

// tips class with field(title and tip)
type Tips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

const (
	empty_string  = ""
	empty_value   = "should not be Empty"
	default_value = "Tips Not Available for Topic"
)

//GetTip returning Tip/Command to the controller
func GetTip(title string, reader readerI) string {
	data, _ := loadTipsFromJson(reader)
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

//getting all tips and titles
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

//unmarshal json data into Tips struct
func loadTipsFromJson(reader readerI) ([]Tips, error) {
	// run an app from main.go -> file path should be "data/tips.json"
	// if want to check all unit test cases ->file path should be "../data/tips.json"
	var path = getJsonFilePath(reader)
	var data []byte
	data, _ = readJsonFile(path, reader)
	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result, nil
}

// getting file path for main file and testing function
func getJsonFilePath(reader readerI) string {
	currentDir, _ := getCurrentWorkingDir(reader)
	// remove base directory from the workingDir when run from test
	baseDir := filepath.Base(currentDir)
	isInTest := os.Getenv("GO_ENV") == "test"
	if isInTest {
		currentDir = strings.ReplaceAll(currentDir, baseDir, "")
	}
	return currentDir + "/data/tips.json"
}

// reading json file data in byte datatypes
func readJsonFile(path string, reader readerI) ([]byte, error) {
	var errFileNotFound = errors.New("failed loading jSON file")
	jsonData, err := reader.readFile(path)
	if err != nil {
		fmt.Println(errFileNotFound)
		return nil, errFileNotFound
	}
	return jsonData, nil
}

//getting current working dir.
func getCurrentWorkingDir(reader readerI) (string, error) {
	var workingDir string
	workingDir, _ = reader.get_wd()
	return workingDir, nil
}

/* INTERFACE IMPLEMENTATIONS FOR MAIN FUNCTION*/

// interface (readerI) implementations for main functionality with two Methods
//"readFile(path string) ([]byte, error)" and "get_wd() (string, error)"
type Reader struct{}

// ReadFile returning file data in bytes and taking file path as a argument
func (reader Reader) readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// get_wd returning working directory
func (reader Reader) get_wd() (string, error) {
	return os.Getwd()
}
