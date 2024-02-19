package helper

import (
	"friday-go/config"
	"net/http"
	twitter "github.com/g8rswimmer/go-twitter/v2"
)



type Authorize struct{
  Token string
}

func (a Authorize) Add(req  *http.Request){
}


func InitAuth(twitterConfig config.TConfig) *twitter.Client{
  client := &twitter.Client{
    Authorizer : Authorize {
      Token: twitterConfig.BearerToken,
    },
    Client : http.DefaultClient,
    Host: "https://api.twitter.com",
  }
  return client
}
