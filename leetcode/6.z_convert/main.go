package main

import (
	"bytes"
	"log"

	//"bufio"
	"fmt"
)

//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

func main() {
	var s string
	//LCIRETOESIIGEDHN
	s = "LEETCODEISHIRING"
	//LDREOEIIECIHNTSG
	s = "LEETCODEISHIRING"
	//"PAHNAPLSIIGYIR"
	s = "PAYPALISHIRING"
	s = "ABC"
	row := 3
	res := convert(s, row)
	fmt.Println(res)
}
func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := new(bytes.Buffer)
	p := 2*numRows - 2

	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	for r := 1; r < numRows-1; r++ {
		res.WriteByte(s[r])

		for k := p; k-r < len(s); k += p {
			res.WriteByte(s[k-r])

			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}

func test1(str string, rowNumber int) string {
	var (
		result string
	)
	arr := make([]string, rowNumber)
	log.Printf("arr:%v", arr)
	return result
}
