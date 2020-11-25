package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var x = 0
	x = 1534236469
	fmt.Println(1 << 31)
	fmt.Println(1e5)
	res := reverse(x)
	fmt.Println(res)
}
func reverse(x int) int {
	var (
		res int
	)
	if x < -1<<31 || x > 1<<31-1 {
		return 0
	}

	return res
}

func reverse2(x int) int {
	var (
		flag = 1
		res  int
		sb   strings.Builder
	)
	//2147483647
	if x > 2147483647 {
		return 0
	}
	if x < -2147483648 {
		return 0
	}
	if x < 0 {
		flag = -1
		x = x * flag
	}
	s := strconv.Itoa(x)
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	s = sb.String()
	//fmt.Println("s:", s)
	res, _ = strconv.Atoi(s)
	if flag == -1 {
		res = res * flag
	}
	if res > 2147483647 {
		return 0
	}
	if res < -2147483648 {
		return 0
	}
	return res
}
