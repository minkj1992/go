# 기본 문법

## 주요 특징

### 간결한 문법

- while 문 x
- 복잡한 if문 대신 switch case에 조건식 첨부


```go

func myFor1() {
	sum := 0
	for i :=0; i<10;i++ {
		sum+=i
	}
	fmt.Println(sum)
		
}

func myFor2() {
	// for문 조건식만
	sum, i := 0, 0
	for i<10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

func myFor3() {
	// for문 조건식 생략
	sum, i := 0, 0
	for {
		if i>=10{
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
		fmt.Printf("%c는 소문자",c)
	case 'A' <= c && c <= 'Z':
		fmt.Printf("%c는 대문자",c)
	}
}

```

### 정적, 동적
- 인터페이스는 덕 타이핑(동적)
- 변수 타입, 인터페이스 덕 타이핑

### 모호함 제거
- 포인터 연산 불허
- `i = i++` 불허
- `++i`, `--i` 불허

### 세미콜론 생략 가능
- 문장의 끝에 go 컴파일러가 세미콜론 자동 삽입
- 두개 이상의 문장 한 줄 작성 시 세미콜론 사용
- **다만 이런 기능 때문에 `{`는 반드시 제어문, 함수 시작 되는 줄의 끝에 써야한다.** 다른 줄에 여는 괄호를 쓰면, 세미콜론이 붙어서 컴파일 에러

## 변수와 상수

### 변수
- var 변수명 타입

```
var a int
var b string
```

- 변수 초기값
  - 이는 변수 값이 할당 전에 gc 할당 되는 것을 방지한다.
```
int : 0
float: 0.0
string(문자열 타입: ""
```

- 변수 여러개 선언