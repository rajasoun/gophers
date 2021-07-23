package client

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cristalhq/oauth2"
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

func FindMapValues(d string) string {
	loadfromEnv()
	for j, i := range values {
		if strings.Contains(j, d) {
			return i
		}
	}
	return ""
}

func getAccessToken() (string, string) {
	config := oauth2.Config{
		ClientID:     FindMapValues("token.client.id"),
		ClientSecret: FindMapValues("token.client.secret"),
		TokenURL:     FindMapValues("token.request.url"),
	}
	// create a client
	new_client := oauth2.NewClient(http.DefaultClient, config)
	token, err := new_client.CredentialsToken(context.Background(), FindMapValues("user.name"), FindMapValues("password"))
	if err != nil {
		fmt.Print(err)
	}
	return token.AccessToken, token.TokenType

}
