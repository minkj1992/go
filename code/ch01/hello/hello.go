package main

import (
  "fmt"
  "log"
  "github.com/minkj1992/go/greetings"
)

func main(){
  log.SetPrefix("greetings: ")
  log.SetFlags(0) // ?

  helloWithGreeting("minkj1992")
  helloWithGreeting("")
}

func helloWithGreeting(name string){
  message, err := greetings.Hello(name)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(message)
}
