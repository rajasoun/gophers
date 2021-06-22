package model

import (
	"encoding/json"
	"errors"

	"io/ioutil"
	"strings"
)

type Tips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

//GetTip returning Tip/Command According to each title
func GetTip(title string) string {
	data, err := loadTipsFromJson("../data/tips.json")
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
func loadTipsFromJson(path string) ([]Tips, error) {
	var errorFile = errors.New("failed loading jSON file")
	data, _ := readJsonFile(path)

	var result []Tips
	json.Unmarshal([]byte(data), &result)
	if string(data) == "" {
		return nil, errorFile
	}
	return result, nil
}
