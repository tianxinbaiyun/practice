@[TOC]( select 块里面的case是随机执行的 )
## 4. 下面代码会触发异常吗？请详细说明

```text
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}


```

### 运行结果

```
panic: hello

1

```


```
panic: hello

goroutine 1 [running]:
main.main()
        /home/develop/go/src/practice/make_study/1005/main.go:18 +0x208

```

## 总结

有可能会发生异常，如果没有selct这段代码，就会出现线程阻塞，当有selct这个语句后，系统会随机抽取一个case进行判断，只有有其中一条语句正常return，此程序将立即执行。