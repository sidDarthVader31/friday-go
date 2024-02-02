package main

import (
	"fmt"
	"friday-go/config"
)
func main(){
  //print twitter config
  fmt.Println("this is init file ")
  fmt.Println("twitter config:", config.GetTwitterConfig())
}
