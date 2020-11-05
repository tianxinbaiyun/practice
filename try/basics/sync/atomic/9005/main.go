package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	for i := int32(0); i < 10; i++ {
		go func() {
			atomicTest(i)
		}()

	}
	time.Sleep(10 * time.Second)
}

func atomicTest(a int32) {
	//var a int32 = 1
	b := atomic.AddInt32(&a, 1)
	fmt.Println(b)
}
