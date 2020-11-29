# go net/http vs gin
> go의 대표적인 http 패키지 `net/http` 패키지와 `gin`을 비교합니다.


## `net/http`
> http 서버를 만들 때 사용되는 기본 라이브러리인 `net/http`

```go
func ListenAndServe(addr string, handler Handler) error // HTTP 연결을 받고, 요청에 응답
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) // 경로별로 요청을 처리할 핸들러 함수를 등록
func Handle(pattern string, handler Handler) // 경로별로 요청을 처리할 핸들러 등록
func FileServer(root FileSystem) Handler // 파일 서버 핸들러
func StripPrefix(prefix string, h Handler) Handler // HTTP 요청 경로에서 특정 문자열을 제거합니다. 예) /assets/hello.js → hello.js
func NewServeMux() *ServeMux // HTTP 요청 멀티플렉서 인스턴스 생성
```

> multiplexing
>> multiplexing은 단순히 말해서 browser가 하나의 connection 상에서 동시에 여러 개의 request를 보내는 기술을 의미한다. Http1.1에서는  connection을 여러개 맺어서 동시에 파일을 받는 것. 하지만 이 경우에도 커넥션끼리의 정보를 주고 받는 문제라던가, 너무 많은 컨텐츠를 소비하는 문제 같은 것들이 발생한다. Http2.0 부터는 동시에 여러 개의 컨텐츠를 받는 것은 동일하지만, queue와 같은 기술을 통해서 하나의 connection만으로 컨텐츠를 동시에 받을 수 있도록 하는 것이다.