# go 환경 세팅
> About go Module

`Go-in-Action`을 하다 소스코드 빌드에서 계속 문제가 발생하였다. GOROOT, GOPATH 등의 환경세팅과 module 세팅이 달라, 정리하고자 한다.


- [go-module-in-action](https://crispgm.com/page/go-module-in-action.html)
- [공식](https://blog.golang.org/using-go-modules)


go 모듈의 주요 동기는 다른 개발자가 작성한 코드 사용 경험을 개선하는 것(종속성 추가)



## go 설치
```bash
$ brew install go
$ go version # go version 확인
$ go version
go version go1.15.3 darwin/amd64
$ go env # 환경 확인
```

- go env 세팅
  - go version 1.15기준으로 설명
```
go env -w GO111MODULE=on
go env -w GOPROXY="https://goproxy.io,direct" # mirror proxy services
```

- 디렉토리 생성/이동
  - 원하는 프로젝트로 이동

- go mod 세팅
  - 의존성들을 가져오도록 명시한다.

```bash
go mod init <원하는 모듈 이름>

```



