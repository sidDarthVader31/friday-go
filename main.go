package main

import (
	"fmt"
	"friday-go/config"
	"friday-go/helper"
	searchstream "friday-go/searchStream"

	"github.com/g8rswimmer/go-twitter/v2"
)
func main(){
  //load env file 


  //print twitter config
  fmt.Println("this is init file ")
  twitterConfig := config.GetTwitterConfig()
  fmt.Println("twitter config:", twitterConfig)

  //initialize twitter 
  twitterClient := helper.InitAuth(twitterConfig)
  fmt.Println("twitter client:", twitterClient)

  //start searchStream 
  rule :=  searchstream.AddRule("#javascript", "javascript")
  ruleArray := [] twitter.TweetSearchStreamRule{rule}
  tweets, err := searchstream.SearchTweet(ruleArray, twitterClient)
  if err == nil{
    fmt.Println("error while searching tweets:", err)
  }

  //retweet tweets 

  for i,v := range tweets {
    tweetId := fmt.Sprintf("%v", v)
    userId := fmt.Sprintf("%v", i)
    go helper.Retweet(&userId, &tweetId, twitterClient)
  }
  fmt.Println("tweets")
  fmt.Println(tweets)
}
