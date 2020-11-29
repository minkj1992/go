# Concurrency in Go

## WaitGroup

- waitGroup 예시
```go
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
    wgCount := 10
    wg.Add(wgCount)
	for i := 0; i < wgCount; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(strconv.Itoa(i) + " goroutine sleeping...")
			time.Sleep(1)
		}(i)
	}

	wg.Wait()   
	fmt.Println("All goroutines complete")
}
```

**waitGroup에서 goroutine의 순서는 random하다.**

- 바람직한 waitGroup **호출 관례**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 10
	var wg sync.WaitGroup

	// Add에 대한 호출은 go루틴 클로저 안에서 실행되면 안된다.
	// 고루틴이 실행되기 전에 wait에 도달할 수도 있기 때문(add 해준 숫자 만큼 wait을 실행한다.)
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i)
	}
	wg.Wait()
}
```