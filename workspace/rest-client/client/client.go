package client

import (
	"bufio"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/cristalhq/oauth2"
)

var values = make(map[string]string)
var openFile = os.Open

func loadDatafromEnv() (map[string]string, error) {
	file, err := openFile("../data.env")
	if err != nil {
		//	fmt.Printf("error opening file: %v\n", err)
		return map[string]string{}, err
	}
	reader := bufio.NewReader(file)
	data, err := readln(reader)
	for err == nil {
		data, err = readln(reader)
	}
	return data, nil
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
func findMapValues(oauthId string) string {
	clientValues, _ := loadDatafromEnv()
	for Key, value := range clientValues {
		if strings.Contains(Key, oauthId) {
			return value
		}
	}
	return "not available"
}

var userName string = findMapValues("user.name")
var pass string = findMapValues("password")

func getAccessToken() (string, string, error) {
	config := oauth2.Config{
		ClientID:     findMapValues("token.client.id"),
		ClientSecret: findMapValues("token.client.secret"),
		TokenURL:     findMapValues("token.request.url"),
	}
	// create a client
	new_client := oauth2.NewClient(http.DefaultClient, config)
	token, err := new_client.CredentialsToken(context.Background(), userName, pass)
	if err != nil {
		//fmt.Println("Erorr from main:", err)
		return "", "", err
	}
	return token.AccessToken, token.TokenType, nil

}
