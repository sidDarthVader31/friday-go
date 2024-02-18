package main

import (
	"fmt"
	"friday-go/config"
  "friday-go/helper"
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


}
