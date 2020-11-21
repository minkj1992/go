# go 환경 세팅
> About go Module

`Go-in-Action`을 하다 소스코드 빌드에서 계속 문제가 발생하였다. GOROOT, GOPATH 등의 환경세팅과 module 세팅이 달라, 정리하고자 한다.

go 모듈의 주요 동기는 다른 개발자가 작성한 코드 사용 경험을 개선하는 것(종속성 추가)

- [go modules wiki](https://github.com/golang/go/wiki/Modules)
- [golang wiki](https://github.com/golang/go/wiki/Modules)
- [golang ref](https://golang.org/ref/mod)


## go 설치

- go 설치
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


## go modules

### Quick Start


- 디렉토리 생성
```bash
mkdir -p <PROJECT_DIR>
cd <PROJECT_DIR>
```

- (Optional vcs)
```bash
git init
git remote add origin <githubrepo>
```

- Init my module

```bash
go mod init github.com/<MYREPO>
```

- Write my go code
- Build and run
```bash
go build -o <MY_MODULE_NAME>
./<MY_MODULE_NAME>
```

- git push
  - `clean`:remove object files and cached files

```bash
go clean # 불필요한 cache, dep, build obj 파일 제거
git push
```

### Q/A

> go mod는 bin의존성을 어떻게 처리하고 있을까?

```bash
$ go env

...

GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="<ROOT_DIR이름>/go"
GOPRIVATE=""
```

go run 실행시, 의존성은 GOPATH에 위치하여 저장되게 된다.

- GOPATH 위치
```bash
.
├── mod
│   ├── cache
│   └── github.com
└── sumdb
    └── sum.golang.org
```

- go.sum
```go
module github.com/minkj1992/test  ----> <MY_MODULE_NAME>

go 1.15

require github.com/webgenie/go-in-action v0.0.0-20160424215153-d6068f2ca89b // indirects

----> <필요한 의존성>
```

**즉 정리하자면, go run에 필요한 bin파일들은 `GOPATH`에 저장되며 실행시 해당 링크를 통해 참조를 실행한다.단 build로 생성된 `실행파일`의 경우 실제로 사용하는 `3rd lib 파일`들을 contains 하는지는 다른 주제로 알아야 할 듯 싶다.**
