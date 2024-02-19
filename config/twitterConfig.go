package config

import (
	"fmt"

	"github.com/spf13/viper"
)
type TConfig struct {
  ConsumerKey string
  ConsumerSecret string
  AccessToken string 
  AccessTokenSecret string
  BearerToken string
}
func GetTwitterConfig() TConfig{
  return getEnvVariables()
}


func getEnvVariables() TConfig{
  viper.SetConfigFile(".env")
  err := viper.ReadInConfig()
  if err!= nil{
    fmt.Println("error while reading config file", err)
  }
  tc := TConfig{
    ConsumerKey: viper.Get("CONSUMER_KEY").(string),
    ConsumerSecret: viper.Get("CONSUMER_SECRET").(string),
    AccessToken: viper.Get("ACCESS_TOKEN").(string),
    AccessTokenSecret: viper.Get("ACCESS_TOKEN_SECRET").(string),
    BearerToken: viper.Get("BEARER_TOKEN").(string),
  }
  return tc
}
