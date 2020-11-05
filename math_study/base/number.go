package main

import (
	"fmt"
	"runtime"
	"time"
)

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func wei(a int) bool {
	fmt.Println(a & 1)
	return false
}

func swap(a, b int) (int, int) {
	//fmt.Printf("start:a:%d,b:%d\n", a, b)
	a ^= b
	//fmt.Printf("a:%d,b:%d\n", a, b)
	b ^= a
	//fmt.Printf("a:%d,b:%d\n", a, b)
	a ^= b
	//fmt.Printf("end:a:%d,b:%d\n", a, b)
	return a, b
}
func swap2(a, b int) (int, int) {
	return b, a
}
func swap3(a, b int) (int, int) {
	tmp := a
	a = b
	b = tmp
	return a, b
}
func main() {
	var (
		//t = time.Now()
		n int = 1e12
	)
	fmt.Println(runtime.GOARCH)
	fmt.Printf("%T\n", n)
	fmt.Println("循环次数:", n)
	go func() {
		var t = time.Now()
		for i := 0; i < n; i++ {
			swap(123, 3567)
		}
		d := time.Since(t)
		fmt.Println("位交换时长:", d)
	}()
	go func() {
		var t = time.Now()
		for i := 0; i < n; i++ {
			swap2(123, 3567)
		}
		d := time.Since(t)
		fmt.Println("多返回值交换时长:", d)
	}()
	go func() {
		var t = time.Now()
		for i := 0; i < n; i++ {
			swap2(123, 3567)
		}
		d := time.Since(t)
		fmt.Println("中间值交换时长:", d)
	}()
	time.Sleep(time.Hour)
}
