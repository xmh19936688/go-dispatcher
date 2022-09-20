# go-dispatcher

go-dispatcher用于将需要处理的数据分发到协程执行，可控制最大并发量。

使用`go-dispatcher`之前：

```go
type T struct{}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan T)

	go dispatch(ch, wg)

	wg.Add(len(list))
	for _, data := range list { ch <- data }
	wg.Wait()
}

func dispatch(ch chan T, wg *sync.WaitGroup) {
	for i := 0; i < maxCurrency; i++ {
		go func() {
			for {
				handler(<-ch)
				wg.Done()
			}
		}()
	}
}

func handler(data T) {
	// do sth with data
}
```

使用`go-dispatcher`之后：

```go
import "github.com/xmh19936688/go-dispatcher/dispatcher"

type T struct{}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan T)

	d := dispatcher.New[T]().MaxCurrency(-1).
		Handler(handler).Chan(ch).WaitGroup(wg)
	go d.Start()

	wg.Add(len(list))
	for _, data := range list {ch <- data}
	wg.Wait()
}

func handler(data T) {
	// do sth with data
}
```
