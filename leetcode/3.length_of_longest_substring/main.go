package main

import (
	"fmt"
	"strings"
)

// 无重复字符的最长子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

func main() {
	str := "zxbsadfhjkadfuewjkszdfkljaD"
	n := lengthOfLongestSubstring(str)
	fmt.Println(n)
}

// lengthOfLongestSubstring
func lengthOfLongestSubstring(s string) int {
	var (
		left   int    = 0
		maxlen int    = 0
		retstr string = ""
	)
	for i := range s {
		if index := strings.Index(retstr, s[i:i+1]); index > -1 {
			left = left + index + 1
		}
		retstr = s[left : i+1]
		if maxlen < len(retstr) {
			maxlen = len(retstr)
		}
	}
	return maxlen
}
