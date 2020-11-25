package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	nums := []int{3, 3}
	//nums := []int{3}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
}
