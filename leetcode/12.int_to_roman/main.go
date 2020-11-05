package main

import (
	"bytes"
	"fmt"
)

//整数转罗马数字
func main() {
	var input int
	input = 20
	output := intToRoman(input)
	fmt.Println(output)
}

func intToRoman(num int) string {
	romanDigit := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanDesc := map[int]string{1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	var buffer bytes.Buffer
	for num != 0 {
		for _, r := range romanDigit {
			if num >= r {
				num -= r
				buffer.WriteString(romanDesc[r])
				break
			}
		}
	}
	return buffer.String()
}
