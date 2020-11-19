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

- switch 옆에 태그의 변수는 어느 자료형이든 쓸 수 있습니다. 
- 태그의 값에 따라 case의 '라벨'과 일치하는 것을 찾고 일치하는 case의 구문을 수행합니다. 
- Go언어에서는 switch 옆에 태그뿐만이 아니라 '표현식'을 쓰는 경우가 있습니다
- if문과 달리 { }의 시작 브레이스를 같은 줄에 쓰지 않아도 실행 됩니다.
- break 따로 입력하지 않아도 해당되는 case만 수행됩니다.

```go
package main

import "fmt"

func main() {
    var num int
    fmt.Print("정수입력: ")
    fmt.Scanln(&num)

    switch num {
        case 0:
          fmt.Println("영")
        case 1:
          fmt.Println("일")
        default:
          fmt.Println("모르겠어요.")
    }
}
```

- 조건식 활용 switch case
```go
package main

import "fmt"

func main() {
    var a, b int

    fmt.Print("Enter int a, b:")
    fmt.Scanln(&a, &b)

    switch {
        case a>b:
            //
        case a<b:
            //    
        default:
          fmt.Println("Don't know")
    }
}
```

```go
package main

import "fmt"

func main() {
    var sel int
    var num1, num2, result float32

    fmt.Scanln(&sel)
    fmt.Scanln(&num1, &num2)

    switch  {
    case condition:
        
    }
}
```

## `break, continue, goto`

#### break
- for, switch, select 사용 가능
- `break 레이블 이름` 이동 가능
  - 빠져나온 블록의 스코프 skip
```go
// 무한 루프로 동작하지 않습니다.
package main

import "fmt"

func main() {
    i:=0
    LABEL1:
        for {
            if i==0{
                break LABEL1
            }
        }
        fmt.Println("END")
}
```

#### continue
- break와 다르게 for 문과 연관해서 사용해야합니다.
  - continue를 만날 경우 해당 반복문의 조건검사 부분으로 다시 이동하기 떄문

#### goto
- 사용하지 마라.
  - break with label과 어떤 차이가 있지.. 진입점 1개만 가지고 있다는 것 말고는 차이가 없어 보이는데

```go
package main

import "fmt"

func main() {
    var num int

    fmt.print("자연수 입력:")
    fmt.Scanln(&num)

    if num == 1 {
        goto ONE
    } else if num ==2 {
        goto TWO
    } else {
        goto DEFAULT
    }

    ONE:
        fmt.Print("1을 입력했습니다.\n")
        goto END
    TWO:
        fmt.Print("2를 입력했습니다.\n")
        goto END
    OTHER:
        fmt.Print("1이나 2를 입력하지 않으셨습니다!\n")
    END:
}
```

- 홀수 구구단
```go
package main

import "fmt"

func main() {
	var i, j int = 2, 1

	for {
		i += 1
		if i == 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		for {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
			j += 1
			if j == i+1 {
				j = 1
				break
			}
		}
		fmt.Println()
	}
}
```

## Array
> go에서는 2개 이상의 변수를 모아 놓은 것을 collection이라 칭함

- go에서 배열은 static(정적, 고정된 크기)
- `var 배열이름 [배열크기]자료형`
- go에서 배열의 크기는 자료형을 구성하는 요소
  - [3]int != [5]int 는 다른 타입으로 처리

- 1차원 배열 선언 방법
```go
package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println(arr1)

	arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, arr1[0], arr1[4])
	// fmt.Println(arr1[-1]) (x)
	// fmt.Println(arr1[1000]) (x)

	arr2 := [4]int{4, 5, 6, 7}
	// var arr2 [4]int = [4]int{4, 5, 6, 7}와 차이가 없을 듯
	arr2[0] = 32
	fmt.Println(arr2)

	var arr3 = [...]int{9, 8, 7, 6, 5} // 자동 배열 크기 조절
	fmt.Println(arr3, len(arr3))

}

```

- 다차원 배열
```go
package main

import "fmt"

func main() {
	var a = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(a[2][2]) // 9
}

```