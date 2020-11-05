package main

import (
	"fmt"
	"time"
)

func main() {
	var n = 50
	t2 := time.Now()
	for i := 1; i < n; i++ {
		r := cal(i)
		fmt.Printf("r:%d,", r)
	}
	fmt.Printf("\napp elapsed:%d\n", time.Since(t2))
	t1 := time.Now()
	for i := 1; i < n; i++ {
		r := Recursion(i)
		fmt.Printf("r:%d,", r)
	}
	fmt.Printf("\napp elapsed:%d\n", time.Since(t1))

}

//Recursion Recursion
func Recursion(n int) int {
	if n <= 2 {
		return 1
	}
	return Recursion(n-1) + Recursion(n-2)
}

func cal(n int) int {
	var (
		result int = 0
		pre    int = 0
		pre2   int = 0
	)
	if n <= 2 {
		return 1
	}
	pre = 1
	pre2 = 1
	for i := 2; i < n; i++ {
		result = pre + pre2
		pre2 = pre
		pre = result

	}

	return result
}
