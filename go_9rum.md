# 한눈에 끝내는 고랭 기초
> [강좌](https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88)

## convention

- go는 camelCase
- ++숫자 못씀
- 3항 연산자가 없다 대신,  if/else syntax is the idiomatic way.

```go
if expr {
    n = trueVal
} else {
    n = falseVal
}

// 또는 inline func 사용
a := func() int { if expr { return 1 } else { return 2 } }()
```

## 형변환
- 파이썬 처럼 간단한 형변환 제공 안하고 있다.
  - int를 bool type으로 전환 따위


## 문자열

- 문자열 선언 방법 2가지
  - ` (Back Quote): Raw String Literal(escape 그대로 표시)
  - ": Interpreterd String literal (escape 동작, 적용된 엔터는 무시)

> ' (Single quotes)의 경우는 byte or rune 타입을 나타낼 때 사용

```go
plusString := "안녕 " + "내 이름은" + "존2세야" 
```

- string 선언한 문자열 타입은 immutable
  - var str string = "hello"와 같이 선언하고 str[2] = 'a'로 수정이 불가능합니다.

## Print rune

Rune is codepoint, thus is a integer.

The following Printf formats work with integer:

- %c → character as is.
- %q → rune syntax. e.g. 'a'.
- %U → Unicode notation. e.g. U+03B1.
- %b → base 2
- %o → base 8
- %d → base 10
- %x → base 16, with lower-case letters for a-f
package main

```go
import "fmt"

// print rune in different formats

func main() {
	var x = '😂'

	fmt.Printf("%c\n", x) // 😂
	fmt.Printf("%q\n", x) // '😂'
	fmt.Printf("%U\n", x) // U+1F602

	fmt.Printf("%b\n", x) // 11111011000000010
	fmt.Printf("%o\n", x) // 373002
	fmt.Printf("%d\n", x) // 128514
	fmt.Printf("%x\n", x) // 1f602
}
```

## variable

- multiple variables
```go
var x, y, x int = 1,1,1
```

## for

- Go언어에서는 while문을 제공하지 않아 for문만 사용할 수 있습니다. 

```go
// while 문 대신 사용 for
package main

import "fmt"

func main(){
    n:=2
    for n<100{
        fmt.Printf("count %d\n",n)
        n*=2
    }
}
```

- 무한 루프
```go
package main

import "fmt"

func main() {
    for {
        fmt.Println("Infinite loop")
    }
}
```

- 3항 다 사용하는 for

```go
 for 초기식; 조건식; 조건 변화식 {
	반복 수행할 구문
}
``` 

```go
package main

import "fmt"

func main() {
    sum:=0

    for i:=1;i<=10;i++ {
        sum += i
    }
}
```

- 초기식 제거 for(2항)

```go
package main

import "fmt"

func main() {
    var num int = 3
    for ; num < 4; num++ {
	fmt.Println(num)
  }
}
```

- `for range`
```go
package main

import "fmt"

func main(){
    var arr [6]int = [6]int{1,2,3,4,5,6}

    // for _, num := range  앞자리 생략 가능
    // for i := range arr 값 생략 가능
    for i, num := range arr {
        fmt.Printf("arr[%d]의 값은 %d 입니다.\n", i, num )
    }
}
```

- 컬렉션을 활용한 for ~ range

```go
package main

import "fmt"

func main() {
    var fruits map[string]string = map[string]string{
        "apple": "red",
        "banana": "yellow",
        "grape": "purplee"
    }

    for fruit, color := range fruits {
        fmt.Printf("%s의 색깔은 %s 입니다. \n", fruit, color)
    }
}
```

## `if/else`

- Go언어에서 조건문은 반드시 Boolean (c와 다르게 1,0 안됨)

- 조건식의 괄호 생략 가능
- 조건문의 중괄호 필수
- 괄호의 시작과 else문은 같은 줄에

```go
if num == 1 {
    //
} else if num == 2 {
    //
} else {
    //
}
```

- 조건식에서 간단한 변수 생성 가능 (`Optional Statement`)
```go
package main

import "fmt"

func main() {
    var num int = 3
    
    if tmpVar:= num * 2; tmpVar == 6 {
	fmt.Printf("%s","HELLO")
	// fmt.Println(tmpVar) 불가능
    }
}
```

- 7과 9의 배수
  - 놀라운게 int를 bool 타입으로 python 처럼 못바꿈..
```go
// 7 또는 9의 배수이면 print 하라( 1~ 100 )
package main

import "fmt"

func main() {
    var num, max_num int = 1, 100
    var isSevenMul, isNineMul int

    for ;num < max_num; num++ {
        isSevenMul = num % 7
        isNineMul = num % 9

        if isSevenMul == 0 || isNineMul ==0 {
            fmt.Printf("%d ", num)
        }else {
            continue
        }
    }
}
```

## switch
> https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/80288/swich%EB%AC%B8%EC%97%90-%EC%9D%98%ED%95%9C-%EC%84%A0%ED%83%9D%EC%A0%81-%EC%8B%A4%ED%96%89