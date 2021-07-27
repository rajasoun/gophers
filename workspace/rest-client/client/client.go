package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"

	"github.com/cristalhq/oauth2"
)

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

var client HTTPClient

func init() {
	client = &http.Client{}
}
func getHttpRequest(header http.Header, url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err)
	}
	request.Header = header
	return client.Do(request)
}

type data struct {
	// todo add more fields
	HasUnlimited bool   `json:"hasUnlimitedLicenses"`
	LastModi     string `json:"lastModifiedBy"`
	Created      string `json:"createdBy"`
}

func getDatafromRestapi(response *http.Response) ([]data, error, string) {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err, ""
	}
	var dataString []data
	json.Unmarshal(bodyBytes, &dataString)
	fmt.Println(len(dataString))
	return dataString, nil, string(bodyBytes)
}

func Run() {
	// 	var configFilepath string = "./configfile" //current working directory
	// 	config, err := loadfromEnv(configFilepath)
	// 	fmt.Println(config)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	token, err := getToken(config)
	// 	fmt.Println(token)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	header := setHeader(token)
	// 	fmt.Println(header)

	// 	req, err := getHttpRequest(header, config.ProductURL)
	// 	fmt.Println(req)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	jsonString, err, _ := getDatafromRestapi(req)
	// 	fmt.Println(jsonString)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
}
