@[TOC](最长公共前缀)

## 最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

```
示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。


```



## 代码
```text
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
		res    string = ""
		flag   int    = 1
		max    int    = 0
		minLen int    = len(strs[0])
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

```


## 运行结果

```
123123
```


## 总结

### 算法

分治

于是，“将整数转换为罗马数字”的过程，就是用上面这张表中右边的数字作为“加法因子”去分解一个整数，目的是“分解的整数个数”尽可能少，因此，对于这道问题，类似于用最少的纸币凑成一个整数，贪心算法的规则如下：

每一步都使用当前较大的罗马数字作为加法因子，最后得到罗马数字表示就是长度最少的。


### 时间复杂度

最坏情况下，我们有 nn 个长度为 mm 的相同字符串。

时间复杂度：O(S)，S 是所有字符串中字符数量的总和，S=m∗n。

时间复杂度的递推式为 T(n)=2⋅T( 2n)+O(m)， 化简后可知其就是 O(S)。
最好情况下，算法会进行 minLen⋅n 次比较，其中 minLen 是数组中最短字符串的长度。

空间复杂度：O(m⋅log(n))

内存开支主要是递归过程中使用的栈空间所消耗的。 一共会进行log(n) 次递归，
每次需要 mm 的空间存储返回结果，所以空间复杂度为 O(m⋅log(n))。











