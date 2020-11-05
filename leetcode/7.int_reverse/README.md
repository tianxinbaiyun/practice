@[TOC](整数反转)

## 整数反转

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。


```
示例 1:
   
   输入: 123
   输出: 321
    示例 2:
   
   输入: -123
   输出: -321
   示例 3:
   
   输入: 120
   输出: 21
   注意:
   

```



## 代码
```text
package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	var x int = 0
	x = 1534236469
	fmt.Println(1 << 31)
	fmt.Println(1e5)
	res := reverse(x)
	fmt.Println(res)
}
func reverse(x int) int {
	var (
		res int = 0
	)
	if x < -1<<31 || x > 1<<31-1 {
		return 0
	}

	return res
}

func reverse2(x int) int {
	var (
		flag int = 1
		res  int = 0
		sb   strings.Builder
	)
	//2147483647
	if x > 2147483647 {
		return 0
	}
	if x < -2147483648 {
		return 0
	}
	if x < 0 {
		flag = -1
		x = x * flag
	}
	s := strconv.Itoa(x)
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	s = sb.String()
	//fmt.Println("s:", s)
	res, _ = strconv.Atoi(s)
	if flag == -1 {
		res = res * flag
	}
	if res > 2147483647 {
		return 0
	}
	if res < -2147483648 {
		return 0
	}
	return res
}

```


## 运行结果
```
2147483648
100000
0
```


## 总结

### 算法

反转整数的方法可以与反转字符串进行类比。

我们想重复“弹出” xx 的最后一位数字，并将它“推入”到 ev 的后面。最后，rev 将与 xx 相反。

要在没有辅助堆栈 / 数组的帮助下 “弹出” 和 “推入” 数字，我们可以使用数学方法。



### 时间复杂度

时间复杂度：O(log(x))，xx 中大约有 log 10(x) 位数字。
空间复杂度：O(1)。







