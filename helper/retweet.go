package helper

import (
	"context"
	"fmt"

	"github.com/g8rswimmer/go-twitter/v2"
)



func Retweet(userId *string, tweetId *string, twitterClient *twitter.Client) {
  _, err := twitterClient.UserRetweet(context.Background(), *userId, *tweetId)
  if err != nil {
    fmt.Printf("error while retweeting tweet with tweet id %v : %v", *tweetId, err)
  }
}
