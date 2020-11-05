package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	input = "CMIX"
	input = "MCMXCIV"
	input = "XX"
	output := romanToInt(input)
	fmt.Println(output)
}
func romanToInt(s string) int {

	ret := 0
	cur := ""
	pre := ""
	for i := range s {
		cur = s[i : i+1]
		if i != 0 {
			pre = s[i-1 : i-1+1]
		}
		switch cur {
		case "I":
			ret++
		case "V":
			ret += 5
			if pre == "I" {
				ret -= 2
			}
		case "X":
			ret += 10
			if pre == "I" {
				ret -= 2
			}
		case "L":
			ret += 50
			if pre == "X" {
				ret -= 20
			}
		case "C":
			ret += 100
			if pre == "X" {
				ret -= 20
			}
		case "D":
			ret += 500
			if pre == "C" {
				ret -= 200
			}
		case "M":
			ret += 1000
			if pre == "C" {
				ret -= 200
			}
		}
	}
	return ret
}
func romanToInt2(s string) int {
	var (
		romanDigit = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		romanDesc  = map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
		res        int
		index      int
	)

	for s != "" {
		for ; index < 13; index++ {
			v := romanDigit[index]
			index := strings.Index(s, v)
			if index == 0 {
				res += romanDesc[v]
				if len(v) == 1 {
					s = s[index+1:]
				} else {
					s = s[index+2:]
				}
				index--
				break
			}
		}
	}
	return res
}
