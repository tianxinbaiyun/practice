package main

import (
	"fmt"
)

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func main() {
	var x int
	x = 225522

	res := isPalindrome(x)

	fmt.Println(res)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var (
		temp = x
		val  = 0
	)
	for x != 0 {
		val = val*10 + x%10
		x = x / 10
	}
	if temp == val {
		return true
	}
	return false
}
