package main

import (
  "fmt"
  "log"
  "github.com/minkj1992/go/greetings"
)

func main(){
  log.SetPrefix("greetings: ")
  log.SetFlags(0) // ?

  hello("minkj1992")
  hello("")
}

func hello(name string){
  message, err := greetings.Hello(name)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(message)
}
