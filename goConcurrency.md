# Concurrency in Go

## WaitGroup

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
	for i := 0; i < wgCount; i++ {
		wg.Add(1)
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

- **waitGroup에서 goroutine의 순서는 random하다.**