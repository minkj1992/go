package greetings

import "fmt"

func Hello(name string) (msg string) {
	msg = fmt.Sprintf("Hi, %v. Welcome!", name)
	return
}