package config

type twitterConfig struct {
  consumerKey string
  consumerSecret string
  accessToken string 
  accessTokenSecret string
}
func GetTwitterConfig() twitterConfig{
  tc := twitterConfig{
  "sdf",
  "sdf",
  "sdf",
  "sdf",
  }
  return tc
}
