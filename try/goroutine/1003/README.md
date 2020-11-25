@[TOC]( gorouter for循环)
## 3. 下面的代码会输出什么，并说明原因

```text
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func init() {
    fmt.Println("Current Go Version:", runtime.Version())
}
func main() {
    runtime.GOMAXPROCS(1)

    count := 10
    wg := sync.WaitGroup{}
    wg.Add(count * 2)
    for i := 0; i < count; i++ {
        go func() {
            fmt.Printf("[%d]", i)
            wg.Done()
        }()
    }
    for i := 0; i < count; i++ {
        go func(i int) {
            fmt.Printf("-%d-", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}


```

### 运行结果

```
Current Go Version: go1.13
-9-[10][10][10][10][10][10][10][10][10][10]-0--1--2--3--4--5--6--7--8-

```

## 总结

两个for循环内部go func 调用参数i的方式是不同的，导致结果完全不同。这也是新手容易遇到的坑。

第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。故go func执行时，i的值始终是10（10次遍历很快完成）。

第二个go func中i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。

