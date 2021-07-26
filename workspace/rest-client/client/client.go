package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cristalhq/oauth2"
	"github.com/spf13/viper"
)

type configuration struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	UserID       string `mapstructure:"user_name"`
	UserPwd      string `mapstructure:"password"`
	TokenURL     string `mapstructure:"token_url"`
	ProductURL   string `mapstructure:"products_url"`
}

var configFilepath string = "../configfile" //current working directory

func LoadfromEnv(path string) (configuration, error) {
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
	fmt.Println(config_data)
	return config_data, nil

}

func getToken() (*oauth2.Token, error) {
	configData, _ := LoadfromEnv(configFilepath)
	oauthConfig := oauth2.Config{
		ClientID:     configData.ClientID,
		ClientSecret: configData.ClientSecret,
		TokenURL:     configData.TokenURL,
	}
	// create a client
	new_client := oauth2.NewClient(http.DefaultClient, oauthConfig)
	token, err := new_client.CredentialsToken(context.Background(), configData.UserID, configData.UserPwd)
	if err != nil {
		return nil, err
	}
	return token, nil
}
