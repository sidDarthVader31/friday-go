package main

import (
	"fmt"
	"friday-go/config"
)
func main(){
  //load env file 


  //print twitter config
  fmt.Println("this is init file ")
  fmt.Println("twitter config:", config.GetTwitterConfig())
}
