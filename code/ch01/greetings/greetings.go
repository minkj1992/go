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

func randomFormat() string {
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v Well met!",
	}

	return formats[rand.Intn(len(formats))]
}


func Hello(name string) (string, error) {
	if name == ""{
		return name, errors.New("empty name")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	// error타입의 nil은 오류가 없음을 의미
	return msg, nil
}

func Hellos(names []string) (map[string]string, error) {
	msgs := make(map[string]string)

	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg

	}
	return msgs, nil

}