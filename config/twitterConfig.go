package config

import (
	"fmt"

	"github.com/spf13/viper"
)
type twitterConfig struct {
  consumerKey string
  consumerSecret string
  accessToken string 
  accessTokenSecret string
}
func GetTwitterConfig() twitterConfig{
  return getEnvVariables()
}


func getEnvVariables() twitterConfig{
  viper.SetConfigFile(".env")
  err := viper.ReadInConfig()
  if err!= nil{
    fmt.Println("error while reading config file", err)
  }
  tc := twitterConfig{
    consumerKey: viper.Get("CONSUMER_KEY").(string),
    consumerSecret: viper.Get("CONSUMER_SECRET").(string),
    accessToken: viper.Get("ACCESS_TOKEN").(string),
    accessTokenSecret: viper.Get("ACCESS_TOKEN_SECRET").(string),
  }
  return tc
}
