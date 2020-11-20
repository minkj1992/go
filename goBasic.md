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

## Slice
- 자료형으로 취급
- 동적으로 크기 변경 가능한 배열
- **배열과 달리 선언만 하면, 배열의 일부분을 가리키는 포인터만 생성** 
  - 즉 실제 배열 메모리 공간은 생성 x
  - 즉 배열 처럼 0으로 자동 초기화 x
  - 이는 동적으로 사이즈가 조절 되니 어느 부분까지 0으로 초기화 해야할지 모르기 때문
- 슬라이스의 초기 값을 지정하지 않고 선언만 한다면 `Nil silce`가 된다..
- make()를 활용하면 선언과 초기화를 동시에 진행 가능
- `[]int{1,2,3,4}`와 같은 식으로 입력하여 값을 초기화하면 새로운 메모리를 할당하면서 그 전의 값은 없어집니다.
- 

- `go/src/runtime/slice.go`
```go
type slice struct {
	array unsafe.Pointer // 포인터
	len   int // 배열 길이
	cap   int // 메모리(배열 생성 시 overflow 체크)
}
```


- 선언 방법
```go
var array [5]int // 배열 자동 초기화 방법
var slicBad []int // (X) 메모리 공간 생성 하지 않습니다.
var sliceGood []int = []int{1, 2, 3, 4} // (O)
```

`sliceGood`은 선언과 동시에 1,2,3,4를 위한 메모리를 만든다는 뜻

#### 슬라이스 복사
배열은 다른 배열의 값을 대입하면, 값 자체가 대입되지만, 슬라이스는 참조 타입이기 때문에 **슬라이스를 복사해온다는 것은 address를 참조한다는 뜻** 입니다.

예를 들면 `l = slice[2:5]`는 slice 에서 2,3,4 인덱스를 참조한다는 뜻입니다.


- 슬라이스 생성 예시
```go
package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10

	fmt.Println(a)

	var b []int

	if b == nil {
		fmt.Println("용량이", cap(b), "길이가", len(b), " Nil Slice입니다.") // 용량이 0 길이가 0  Nil Slice입니다.
	}
}

```

#### make() 이용한 슬라이스 선언
> 선언만 해도 초기화 가능한 방법 make()
-  `make(슬라이스 타입, 슬라이스 길이, 슬라이스의 용량)`
   -  cap 생략 가능, 이때는 cap이 len과 같은 값으로 선언 됩니다.
   - 모든 요소가 0인 슬라이스 생성


#### append()

```go
package main

import "fmt"

func main() {
	s := make([]int, 0, 3)

	for i := 1; i <= 10; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
	fmt.Println(s)
}

```

- 확인해본다면 cap은 double로 사이즈를 늘려나간다.
- 좀 더 구체적으로는 "용량이 초과하는 경우에는 설정한 용량만큼 새로운 배열을 생성하고 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당하는 방식" 입니다.

- `append()`
```go
package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)

	fmt.Println(sliceA) // [1 2 3 4 5 6]
	fmt.Println(sliceB) // [4 5 6]
}
```

- `copy()`
```go
package main

import "fmt"

func main() {
	sliceA := []int{0, 1, 2}
	sliceB := make([]int, len(sliceA), cap(sliceA)*2)

	copy(sliceB, sliceA)              // A를 B에 붙인다.
	fmt.Println(sliceB)               // standart lib
	println(len(sliceB), cap(sliceB)) // 디버깅용 사용, fmt 쓰는게 맞다.

}
```

- shallowcopy vs deepcopy

```go
package main

import "fmt"

func main() {
	c := make([]int, 0, 3) //용량이 3이고 길이가0인 정수형 슬라이스 선언
	c = append(c, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(c), cap(c)) // 7,8

	shallow := c[1:3] //인덱스 1요소부터 2요소까지 복사
	shallow[0] = -1
	fmt.Println(shallow) // [-1 3]

	fmt.Println(c) //슬라이스 l의 값을 바꿨는데 c의 값도 바뀜
	//값을 복사해온 것이 아니라 기존 슬라이스 주솟값을 참조
	// [1 -1 3 4 5 6 7]

	deep := make([]int, len(c))
	copy(deep, c)

	deep[0] = -1
	fmt.Println(c, deep) // [1 -1 3 4 5 6 7] [-1 -1 3 4 5 6 7]
}

```

## map

- 슬라이스와 맵의 공통점은 두 컬랙션 모두 값을 직접적으로 저장하는 것이 아닌 `참조 타입(Reference type)`
- `var 맵이름 map[key자료형]value자료형`
  - ex) `var a map[int]string`로 선언 하면 `Nil map`생성

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var m1 map[int]string

	fmt.Println(m1)                 // map[]
	fmt.Println(reflect.TypeOf(m1)) // map[int]string
	fmt.Println(m1 == nil)          // true

	var m2 = map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}

	fmt.Println(m2, len(m2))        // map[a:a b:b c:c] 3
	fmt.Println(reflect.TypeOf(m2)) // map[string]string
}

```

#### map add, update, delete

- add: `m[key]=value`
- update: `m[key]=newValue`
- delete: `delete(m, key)`

#### map 읽기

- index based collection들은(array, slice) idx 넘어가면 에러
- 그러나, map의 경우 key 유무에 따라 true/false 값을 반환합니다. (자료형에 따라 0이나 "")

- 동작 방식
```go
package main

import "fmt"

func main() {
	var m = make(map[int]int)
	m[1] = 1

	fmt.Println(m[1]) // 1
	fmt.Println(m[2]) // 0
	// fmt.Println(m["삼"]) // type untyped string

	v1, isExist1 := m[1] // value, bool
	v2, isExist2 := m[2]

	fmt.Println(v1, isExist1) // 1 true
	fmt.Println(v2, isExist2) // 0 false

	v1, _ = m[1]
	_, isExist1 = m[1]
}
```

## func

- `func 함수이름 (변수이름 변수타입) 반환형`
  - ex) `func f1 (a int) int {}`
- '반환값'이 여러 개일 수 있다. 이럴 때는 '반환형'을 괄호로 묶어 개수만큼 입력해야한다. `(반환형1, 반환형2)`형식, **두 형이 같더라도 두 번 써야 한다**
- **블록 시작 브레이스({)가 함수 선언과 동시에 첫 줄에 있어야 한다**(모든 용법을 이렇게 쓰는 것이 좋습니다).

```go
package main

import "fmt"

func guide() {
	fmt.Println("Type 2 int with space:")
}

func input() (int, int) {
	var a, b int
	fmt.Scanln(&a, &b)
	return a, b
}

func multi(a, b int) int {
	return a * b
}

func printResult(num int) {
	fmt.Printf("Answer is %d\n", num)
}

func main() {
	guide()
	num1, num2 := input()
	result := multi(num1, num2)
	printResult(result)
}
```

## Pass by reference
> 매개변수

#### C vs Go
- C언어에서는 배열이름 자체가 배열의 첫번째 인덱스 요소의 주솟값인데 Go언어는 그런 것이 없습니다. 
  - go에서는 **주솟값은 어떤 변수 앞에 &를 붙이는 것** 정도 기억
- C언어에서는 "*(배열이름+인덱스)"는 "배열이름[인덱스]"와 같은 기능을 했는데 Go언어는 그런 것이 없습니다. **직접 참조를 원하면 포인터 변수 앞에 *를 붙이**는 것만 기억
- pass by ref로 전달된 address value는 *로 직접참조 해야 사용 가능

```go
package main

import "fmt"

func printSquare(a *int) {
	// ref의 값 변경
	*a *= *a
}

func main() {
	a := 4
	printSquare(&a)
	fmt.Println(a) //16

}
```

## variable argument (가변 인자(= 가변 인수))
> '가변 인자 함수'는 전달하는 매개변수의 개수를 고정한 함수가 아니라 함수를 호출할 때마다 매개변수의 개수를 다르게 전달할 수 있도록 만든 함수

- Go언어의 가변 인자 함수는 동일한 형의 매개변수를 n개 전달 가능
- n개의 동일한 형의 매개변수를 전달
- 전달된 변수들은 `슬라이스 타입`
- `func 함수이름(매개변수이름 ...매개변수형) 반환형`
  - ex) `func f(nums ...int) int {}`

```go
package main

import "fmt"

func sum(nums ...int) int {
	var result int
	for _, v := range nums {
		result += v
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...))
}

```

## Named Return Parameter
> 이름이 붙여진 반환 값

- 여러 개의 값을 반환할때 타입이 다양하면 가독성이 좋지 못하다.
- 따라서 Named return parameter를 사용

- 특징
  - (반환값이름1 반환형1, 반환값이름2 반환형2, 반환값이름3 반환형3, ...)
    - 주의 할 점은 반환형 자체가 변수 선언입니다.
  - `return`은 반드시 명시해주어야 한다.
  - 반환 값이 하나라도 명시한다면 괄호 안에 써주어야 합니다.

```go
package main

import "fmt"

func memberList(names ...string) (count int, list []string) {
	for _, name := range names {
		list = append(list, name)
		count++
	}
	return
}

func inputMember() (list []string) {
	for {
		var name string

		fmt.Println("이름을 입력하시오:")
		fmt.Scanln(&name)

		if name == "0" {
			break
		} else {
			list = append(list, name)
		}
	}
	return
}

func main() {
	count, list := memberList(inputMember()...) //함수를 변수처럼 사용할 수 있습니다

	fmt.Println(count, list) // 6 [john peter kayden leoo victoria martin]
}
```

## Anonymous func(익명함수)

- 이름있는 함수의 단점: **프로그램 속도 저하**
  - 함수 선언 자체가 프로그래밍 전역으로 초기화되면서 메모리를 잡아먹기 때문입니다.
  - 기능을 수행할 때마다 함수를 찾아서 호출해야하기 때문입니다.

```go
package main

import "fmt"

func addDeclared(nums ...int) (result int) {
	for _, v := range nums {
		result += v
	}
	return
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}

	addAnonymous := func(nums ...int) (result int) {
		for _, v := range nums {
			result += v
		}
		return
	}

	fmt.Println(addDeclared(nums...)) // 28
	fmt.Println(addAnonymous(nums...)) // 28
}
```

- 결과는 같지만 선언 함수와 익명함수는 내부적으로 읽는 순서가 다릅니다.
- 선언함수는 프로그램 시작과 동시에 읽어들이지만, 익명함수는 자신의 자리에서 실행되기 때문에 실제 선언된 곳에서 읽혀집니다. (전역변수와 지역변수)

## First-Class Func (일급 함수)

- 일급함수
  - 함수를 args(인자)로 전달
  - 함수를 변수에 할당
  - 함수를 reuturn



```go
package main

import "fmt"

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	multi := func(i int, j int) int {
		return i * j
	}

	r1 := calc(multi, 10, 20) //200
	r2 := calc(func(x int, y int) int { return x + y }, 10, 20) // 30

	fmt.Println(r1)
	fmt.Println(r2)
}
```

## type문을 사용한 함수 원형 정의
> 함수를 인자로 사용할때 args type, return type이 늘어나면 가독성이 떨어진다. 이를 해결

- 함수의 원형 type을 만들어 가독성을 늘리자

```go
package main

import "fmt"
import s "strings"

type numCalculator func([]int) int
type strCalculator func([]string) string

func calNum(f numCalculator, nums ...int) int {
	result := f(nums)
	return result
}

func calStr(f strCalculator, str ...string) string {
	result := f(str)
	return result
}

func main() {
	sum := func(nums ...int) (total int) {
		for _, v := range nums {
			total += v
		}
		return
	}

	dupStr := func(str ...string) (total string) {
		for _, v := range str {
			total += v
		}
		return s.Repeat(total, 2)
	}

	nums := [5]int{1, 2, 3, 4, 5}
	r1 := calNum(sum, nums...)                    // TODO: 에러
	r2 := calStr(dupStr, "a", "b", "c", "d", "e") // TODO: 에러

	fmt.Println(r1)
	fmt.Println(r2)
}

```

## Closure
> https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/81717/%EC%99%B8%EB%B6%80-%EB%B3%80%EC%88%98-%EC%A0%91%EA%B7%BC-%ED%81%B4%EB%A1%9C%EC%A0%80