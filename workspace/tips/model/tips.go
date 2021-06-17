package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Tips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

var ErrInsufficient = errors.New("file not found")
var path = "../data/tips.json"

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
func readJsonFile(path string) ([]byte, error) {
	absPath, _ := filepath.Abs(path)
	jsonData, err := ioutil.ReadFile(absPath)
	if err != nil {
		//log.Fatal(err)
		return nil, ErrInsufficient
	}
	return jsonData, nil
}

//loading json data into Tips struct
func LoadTipsFromJson() ([]Tips, string) {
	data, _ := readJsonFile(path)
	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result, string(data)
}

// to do
// assertError := func(t testing.TB, got error, want string) {
//     t.Helper()
//     if got == nil {
//         t.Fatal("didn't get an error but wanted one")
//     }

//     if got.Error() != want {
//         t.Errorf("got %q, want %q", got, want)
//     }
// }
