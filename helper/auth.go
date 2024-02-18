package helper

import (
	"flag"
	"fmt"
	"friday-go/config"
	"net/http"
	twitter "github.com/g8rswimmer/go-twitter/v2"
)



type Authorize struct{
  Token string
}

func (a Authorize) Add(req  *http.Request){
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}


func InitAuth(twitterConfig config.TConfig) *twitter.Client{
  token := flag.String("token", "", twitterConfig.AccessToken)
  client := &twitter.Client{
    Authorizer : Authorize {
      Token: *token,
    },
    Client : http.DefaultClient,
    Host: "https://api.twitter.com",
  }
  return client
}
