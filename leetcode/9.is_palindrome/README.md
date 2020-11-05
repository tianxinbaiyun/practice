@[TOC](回文数)

## 回文数

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。


```
示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

```



## 代码
```text
package main

import (
	"fmt"
)

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func main() {
	var x int
	x = 225522

	res := isPalindrome(x)

	fmt.Println(res)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var (
		temp int = x
		val  int = 0
	)
	for x != 0 {
		val = val*10 + x%10
		x = x / 10
	}
	if temp == val {
		return true
	}
	return false
}

```


## 运行结果
```
true

```


## 总结

### 算法

反转整数的方法可以与反转字符串进行类比。

我们想重复“弹出” xx 的最后一位数字，并将它“推入”到 ev 的后面。最后，rev 将与 xx 相反。

要在没有辅助堆栈 / 数组的帮助下 “弹出” 和 “推入” 数字，我们可以使用数学方法。




### 时间复杂度

时间复杂度：O(log 10(n))，对于每次迭代，我们会将输入除以10，因此时间复杂度为 O(log 10(n))。

空间复杂度：O(1)。








