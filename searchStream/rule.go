package searchstream

import (
	"context"
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
func SearchTweet(streamRule []twitter.TweetSearchStreamRule, client *twitter.Client) (*twitter.TweetSearchStreamAddRuleResponse, error){
  searchStreamRule, err := client.TweetSearchStreamAddRule(context.Background(), streamRule, true)
  if err != nil {
    fmt.Printf("error while adding rules: %v :: %v", streamRule, err)
    return nil, err
  }

  return searchStreamRule, nil

}
