package searchstream

import (
	"context"
	"encoding/json"
	"fmt"

  "github.com/g8rswimmer/go-twitter/v2"
)




func AddRule(value string, tag string) twitter.TweetSearchStreamRule{
  streamRule := twitter.TweetSearchStreamRule{
    Value : value,
    Tag : tag,
  }
  return streamRule
}

//(*TweetSearchStreamAddRuleResponse, error) 
func SearchTweet(streamRule []twitter.TweetSearchStreamRule, client *twitter.Client) ([]byte, error){
  fmt.Println("stream rule received:", streamRule)
  searchStreamRule, err := client.TweetSearchStreamAddRule(context.Background(), streamRule, false)
  fmt.Println("search stream ruile:", searchStreamRule)
  if err != nil {
    fmt.Printf("error while adding rules: %v :: %v", streamRule, err)
    return nil, err
  }

  enc, err := json.Marshal(searchStreamRule)
  if err != nil{
    fmt.Printf("error while parsing search tweets %v", err)
  }

  return enc, nil

}
