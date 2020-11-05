package main

import (
	"fmt"
	"strconv"
	"strings"
)

//请你来实现一个 atoi 函数，使其能将字符串转换成整数。
func main() {
	var s string = ""
	s = "  -11111111111111111111111419with words"
	//s = "-2147483649"
	s = "--5+"
	//s = "  3.14159"
	//s = "words and 987"
	res := myAtoi(s)
	fmt.Println(res)
	//fmt.Println('+' - '0')
}
func myAtoi(str string) int {
	var (
		res    int
		maxLen int
		left   int
		right  int
		index  int
		s      string
	)
	str = strings.TrimSpace(str)
	for k, v := range str {
		c := v - '0'
		if (c >= 0 && c < 10) || ((c == -3 || c == -5) && maxLen <= 1) {
			right = k + 1
		} else {
			break
		}
		if maxLen < right-left {
			maxLen = right - left
			index = left
		}
		//fmt.Println("v:", c, "l:", left, "r:", right, "index:", index)
	}

	s = str[index : index+maxLen]
	//fmt.Println("s:", s, "index:", index, "max:", maxLen)
	res, _ = strconv.Atoi(s)
	if res >= 1<<31 {
		return 1<<31 - 1
	}
	if res < -1<<31 {
		return -1 << 31
	}
	return res
}
