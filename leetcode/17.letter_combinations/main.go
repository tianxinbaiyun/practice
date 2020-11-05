package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	res := letterCombinations("45678")
	fmt.Println(time.Since(t))
	fmt.Println(res)
}
func letterCombinations(digits string) []string {
	var (
		chars map[int32][]string
		res   []string
	)

	chars = map[int32][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	for k, v := range digits {
		if _, ok := chars[v]; !ok {
			return []string{}
		}
		if k == 0 {
			res = chars[v]
		} else {
			tmp := make([]string, 0)
			for _, vv := range res {
				for _, vvv := range chars[v] {
					tmp = append(tmp, fmt.Sprintf("%s%s", vv, vvv))
				}
			}
			res = tmp
		}
	}
	return res
}

func letterCombinations0(digits string) []string {
	var (
		chars map[string][]string
		res   []string
	)

	chars = map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	for k, v := range digits {
		vStr := string(v)
		if _, ok := chars[vStr]; !ok {
			return []string{}
		}
		if k == 0 {
			res = chars[vStr]
		} else {
			tmp := make([]string, 0)
			for _, vv := range res {
				for _, vvv := range chars[vStr] {
					tmp = append(tmp, fmt.Sprintf("%s%s", vv, vvv))
				}
			}
			res = tmp
		}
	}
	return res
}
