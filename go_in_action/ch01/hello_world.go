package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	myFor1()
	myFor2()
	myFor3()
	mySwitch()
}

func myFor1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}

func myFor2() {
	// for문 조건식만
	sum, i := 0, 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

func myFor3() {
	// for문 조건식 생략
	sum, i := 0, 0
	for {
		if i >= 10 {
			break
		}
		sum += i
		i++
	}
	fmt.Println(sum)
}

func mySwitch() {
	// case에 조건식 사용
	c := 'a'
	switch {
	case '0' <= c && c <= '9':
		fmt.Printf("%c는 숫자", c)
	case 'a' <= c && c <= 'z':
		fmt.Printf("%c는 소문자", c)
	case 'A' <= c && c <= 'Z':
		fmt.Printf("%c는 대문자", c)
	}
}
