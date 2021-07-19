package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var values = make(map[string]string)

func loadfromEnv() map[string]string {
	file, err := os.Open("data.env")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}
	reader := bufio.NewReader(file)
	data, err := readln(reader)
	for err == nil {
		data, err = readln(reader)
	}
	return data

}

func readln(reader *bufio.Reader) (map[string]string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line     []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = reader.ReadLine()
		lineBy := string(line)
		if strings.Contains(lineBy, "=") {
			res1 := strings.Index(lineBy, "=")
			values[lineBy[0:(res1)]] = lineBy[res1+1:]
		}
	}
	return values, err
}

// Config describes a 3-legged OAuth2 flow.
// type Config struct {
// 	ClientID     string
// 	ClientSecret string
// 	UserName     string
// 	Password     string
// 	TokenURL     string
// 	ProductURL   string
// 	ModifiedURL  string
// 	CreatedURL   string
// }
