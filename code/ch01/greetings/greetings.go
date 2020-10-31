package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	if name == ""{
		return name, errors.New("empty name")
	}

	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	// error타입의 nil은 오류가 없음을 의미
	return msg, nil
}