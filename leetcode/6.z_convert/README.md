@[TOC](Z 字形变换)

## Z 字形变换

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。


```
比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G


```



## 代码
```text
package main

import (
	"bytes"
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


```


## 运行结果
```
ABC

```


## 总结

### 算法

按照与逐行读取 Z 字形图案相同的顺序访问字符串。

首先访问 行 0 中的所有字符，接着访问 行 1，然后 行 2，依此类推...

对于所有整数 k，

行 0 中的字符位于索引 k(2⋅numRows−2) 处;
行 numRows−1 中的字符位于索引 k(2⋅numRows−2)+numRows−1 处;
内部的 行 ii 中的字符位于索引 k(2⋅numRows−2)+i 以及 (k+1)(2⋅numRows−2)−i 处;



### 时间复杂度

时间复杂度：O(n)，其中 n == (s)n==len(s)。每个索引被访问一次。
空间复杂度：O(n)。对于 C++ 实现，如果返回字符串不被视为额外空间，则复杂度为 O(1)。







