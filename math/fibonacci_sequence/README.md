@[TOC](斐波那契数列)

## 斐波那契数列

斐波那契数列（Fibonacci sequence），
又称黄金分割数列、因数学家列昂纳多·斐波那契（Leonardoda Fibonacci）以兔子繁殖为例子而引入，
故又称为“兔子数列”，指的是这样一个数列：1、1、2、3、5、8、13、21、34、……

在数学上，斐波那契数列以如下被以递推的方法定义：F(1)=1，F(2)=1, F(n)=F(n - 1)+F(n - 2)（n ≥ 3，n ∈ N*）在现代物理、准晶体结构、化学等领域，


## 实现代码

```text

package main

import (
	"fmt"
	"time"
)
func main() {
	var n int = 20
	t1:=time.Now()
	for i := 1; i < n; i++ {
		r := Recursion(i)
		fmt.Printf("r:%d,", r)
	}
	fmt.Printf("\napp elapsed:%d\n", time.Since(t1))
	t2:=time.Now()
	for i := 1; i < n; i++ {
		r := cal(i)
		fmt.Printf("r:%d,", r)
	}
	fmt.Printf("\napp elapsed:%d\n", time.Since(t2))
}

func Recursion(n int) int {
	if n <= 2 {
		return 1
	}
	return Recursion(n-1) + Recursion(n-2)
}

func cal(n int) int {
	var (
		result int
		pre    int
		pre2   int
	)
	if n <= 2 {
		return 1
	}
	pre = 1
	pre2 =1
	for i := 2; i < n; i++ {
		result = pre + pre2
		pre2 = pre
		pre = result

	}

	return result
}

```

## 运行结果


```
r:1,r:1,r:2,r:3,r:5,r:8,r:13,r:21,r:34,r:55,r:89,r:144,r:233,r:377,r:610,r:987,r:1597,r:2584,r:4181,
app elapsed:127404
r:1,r:1,r:2,r:3,r:5,r:8,r:13,r:21,r:34,r:55,r:89,r:144,r:233,r:377,r:610,r:987,r:1597,r:2584,r:4181,
app elapsed:42639

```

## 总结

1.运行效率:使用递归运算的效率比使用递加运算效率低

2.随着n的增加,递归效率变得更差,例如:

当n=10,递归运算所用时间与递增运算所用时间相差不大,

当n=20,递归运算所用时间是递增运算所用时间的1.5~4倍,

当n=30,递归运算所用时间是递增运算所用时间的100~150倍,

当n=40,递归运算所用时间是递增运算所用时间的3,000~10,000多倍,

当n=50,递归运算所用时间是递增运算所用时间的400,000~100,000,000多倍,