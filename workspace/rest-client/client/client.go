package client

import (
	"context"
	"fmt"
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
	//client := &http.Client{}
	//config, _ := loadfromEnv(configFilepath)
	//token, _ := getToken()
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err)
	}
	request.Header = header
	return client.Do(request)
}

// func CallfromMain() {
//var configFilepath string = "../configfile" //current working directory
// 	config, _ := loadfromEnv(configFilepath)
// 	token, _ := getToken(config)
// 	header := setHeader(token)
// 	getHttpRequest(header, config.ProductURL)
// }
