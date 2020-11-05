package main

import (
	"fmt"
	//"strings"
)

func main() {
	res := StringToInt("+12345A")
	fmt.Println(res)

	//PrintSort("asd")
}

// StringToInt
//问题1：实现一个函数stringToInt,实现把字符串转换成整数这个功能，
//不能使用atoi或者其他类似的库函数。
//注意各种异常边界
//（和错误判断处理比如正负号，非法字符如非数字型字符,
//遇到首个非法字符即截断返回）。
//举例：输入“+12345A”输出12345
func StringToInt(str string) int {
	var (
		res  int
		flag bool
	)
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			tmp := str[i] - '0'
			res = res*10 + int(tmp)
		} else if (i == 0 && str[i] == '-') || (i == 0 && str[i] == '+') {
			if str[i] == '-' {
				flag = true
			}
			continue
		} else {
			break
		}
	}
	if flag {
		res = -1 * res
	}
	return res
}

//问题2：输入一个字符串(字符都不一样)，
//打印出该字符串中字符的所有排列。
//例如输入字符串abc。
//则打印出由字符a、b、c
//所能排列出来的所有字符串abc、acb、bac 、bca、cab 和cba
func PrintSort(str string) {
	//strMap := make(map[int]string)
	//sLen := len(str)
	//res := ""
	//for k, v := range str {
	//	strMap[k] = fmt.Sprintf("%c", v)
	//}
	//fmt.Println(strMap)

	//for i := 0; i < sLen; i++ {
	//	res = fmt.Sprintf("%c", str[i])
	//	j := i + 1
	//	for {
	//		if len(res) > sLen {
	//			break
	//		} else if len(res) < sLen {
	//			if strings.Index(res, strMap[j]) < -1 {
	//				res = fmt.Sprintf("%s%s", res, strMap[i])
	//			}
	//			j = j + 1
	//		} else {
	//			fmt.Println(res)
	//		}
	//
	//	}
	//}
}
