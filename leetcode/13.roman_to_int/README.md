@[TOC](罗马数字转整数)

## 罗马数字转整数

罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

| 字符 | 数值 | 
|---|---|
|I|1|
|V|5|
|X|10|
|L|50|
|C|100|
|D|500|
|M|1000|

例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。




```
示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.



```



## 代码
```text
package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string = ""
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
	for i, _ := range s {
		cur = s[i : i+1]
		if i != 0 {
			pre = s[i-1 : i-1+1]
		}
		switch cur {
		case "I":
			ret += 1
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
		romanDigit     = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		romanDesc      = map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
		res        int = 0
		index      int = 0
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

```


## 运行结果

```
20
```


## 总结

### 算法

贪心算法

于是，“将整数转换为罗马数字”的过程，就是用上面这张表中右边的数字作为“加法因子”去分解一个整数，目的是“分解的整数个数”尽可能少，因此，对于这道问题，类似于用最少的纸币凑成一个整数，贪心算法的规则如下：

每一步都使用当前较大的罗马数字作为加法因子，最后得到罗马数字表示就是长度最少的。


### 时间复杂度

复杂度分析：

时间复杂度：O(1)，虽然看起来是两层循环，但是外层循环的次数最多 12，内层循环的此时其实也是有限次的，综合一下，时间复杂度是 O(1)。

空间复杂度：O(1)，这里使用了两个辅助数字，空间都为 13，还有常数个变量，故空间复杂度是 O(1)。








