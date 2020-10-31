package main

import (
  "fmt"
  "github.com/minkj1992/go/greetings"
)

func main(){
  fmt.Println("Hello, World!")

  message := greetings.Hello("Minwook")
  fmt.Println(message)
}
