package main

import "fmt"

func main() {
	var strs []string
	strs = []string{"flower456", "flowerzcx", "flower123123"}
	strs = []string{"123123"}
	output := longestCommonPrefix(strs)
	fmt.Println(output)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var (
		res    = ""
		flag   = 1
		max    = 0
		minLen = len(strs[0])
	)
	tmp := make(map[int]string, 0)
	for k, v := range strs {
		tmp[k] = v
		if minLen > len(tmp[k]) {
			minLen = len(tmp[k])
		}
	}
	for i := 0; i < minLen; i++ {
		for _, v := range tmp {
			if v[i] != tmp[0][i] {
				flag = 0
				break
			}
		}
		if flag == 0 {
			break
		}
		max++
		flag = 1
	}
	res = tmp[0][0:max]
	return res
}
