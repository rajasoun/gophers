package model

import (
	"encoding/json"
	"io/ioutil"
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
		//fmt.Println(data)
		//if strings.Compare(title, data[index].Title) == 0 {
		//if title == data[index].Title {
		//	return data[index].Tip
		//}
		if strings.Contains(data[index].Title, title) {
			return data[index].Tip
		}
	}
	return "Tips Not Available for Topic"
}

//reading json data from file
func readJsonFile() ([]byte, error) {
	jsonData, _ := ioutil.ReadFile("data/allJson.json")

	//jsonData, _ := ioutil.ReadFile("data/tips.json")
	return jsonData, nil
}

//loading json data into Tips struct
func LoadTipsFromJson() ([]Tips, string) {
	data, _ := readJsonFile()
	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result, string(data)
}
