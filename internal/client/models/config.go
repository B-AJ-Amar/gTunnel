package models

type ClientConfig struct {
	AccessToken string `mapstructure:"access_token"`
	ServerURL   string `mapstructure:"server_url"`
}
