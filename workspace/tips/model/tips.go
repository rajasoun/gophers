package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Tips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

//GetTip returning Tip/Command According to each title
func GetTip(title string) string {
	data, _ := LoadTipsFromJson()
	commands := GetAllCommands(data, title)
	for _, tip := range commands {
		return tip
	}
	fmt.Println(" \n ")
	return "Tips Not Available for Topic"
}

func GetAllCommands(data []Tips, title string) []string {
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
func ReadJsonFile(path string) ([]byte, error) {
	var ErrInsufficient = errors.New("file not found")
	absPath, _ := filepath.Abs(path)
	jsonData, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, ErrInsufficient
	}
	return jsonData, nil
}

//loading json data into Tips struct
func LoadTipsFromJson() ([]Tips, string) {
	var path = "../data/tips.json"
	data, _ := ReadJsonFile(path)
	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result, string(data)
}
