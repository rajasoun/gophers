package model

import (
	"encoding/json"
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
	for index := range data {
		if strings.Contains(data[index].Tip, title) {
			return data[index].Title + " : " + data[index].Tip
		}
	}
	return "Tips Not Available for Topic"
}

//reading json data from file
func readJsonFile() ([]byte, error) {
	absPath, _ := filepath.Abs("../data/tips.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	jsonData, _ := ioutil.ReadFile(absPath)
	return jsonData, nil
}

//loading json data into Tips struct
func LoadTipsFromJson() ([]Tips, string) {
	data, _ := readJsonFile()
	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result, string(data)
}

//to do add error code and test cases
