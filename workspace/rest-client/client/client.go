package client

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"

	"github.com/cristalhq/oauth2"
)

var (
	client      HTTPClient
	httpRequest = http.NewRequest
	read_All    = io.ReadAll
)

func init() {
	client = &http.Client{}
}

type configuration struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	UserID       string `mapstructure:"user_name"`
	UserPwd      string `mapstructure:"password"`
	TokenURL     string `mapstructure:"token_url"`
	ProductURL   string `mapstructure:"products_url"`
}

func loadfromEnv(path string) (configuration, error) {
	var config_data configuration
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return configuration{}, err
	}
	viper.Unmarshal(&config_data)
	return config_data, nil

}

func getToken(configData configuration) (*oauth2.Token, error) {
	oauthConfig := oauth2.Config{
		ClientID:     configData.ClientID,
		ClientSecret: configData.ClientSecret,
		TokenURL:     configData.TokenURL,
	}
	// create a client
	//to do mock the CredentialsToken
	new_client := oauth2.NewClient(http.DefaultClient, oauthConfig)
	token, err := new_client.CredentialsToken(context.Background(), configData.UserID, configData.UserPwd)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func setHeader(token *oauth2.Token) http.Header {
	value := token.TokenType + " " + token.AccessToken
	header := http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{value},
	}
	return header
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func getHttpRequest(header http.Header, url string) (*http.Response, error) {
	request, err := httpRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	request.Header = header
	return client.Do(request)
}

type data struct {
	// todo add more fields
	//todo put in model package
	HasUnlimited bool   `json:"hasUnlimitedLicenses"`
	LastModi     string `json:"lastModifiedBy"`
	Created      string `json:"createdBy"`
}

func getDatafromRestapi(response *http.Response) ([]data, string, error) {
	bodyBytes, err := read_All(response.Body)
	if err != nil {
		return nil, "invalid data", err
	}
	var dataString []data

	json.Unmarshal(bodyBytes, &dataString)
	return dataString, string(bodyBytes), nil
}

func writeDataintoJson(jsonData []data, fileSavePath string) error {
	json_Data, _ := json.MarshalIndent(jsonData, "", "")
	//var fileSavePath ="configfile/configJsonData.json"
	err := ioutil.WriteFile(fileSavePath, json_Data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func writeintoDatabase() error {
	return nil

}
func Run() {

}
