package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


func Hello(name string) (string, error) {
	if name == ""{
		return name, errors.New("empty name")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	// error타입의 nil은 오류가 없음을 의미
	return msg, nil
}

func randomFormat() string {
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v Weel met!",
	}

	return formats[rand.Intn(len(formats))]
}