package restclient

import (
	"os"
)

//var values = make(map[string]string)

var fileRead = os.ReadFile
var path = "data.env"

func loadfromEnv() string {
	data, err := fileRead(path)
	if err != nil {
		return ""
	}
	return string(data)

}
