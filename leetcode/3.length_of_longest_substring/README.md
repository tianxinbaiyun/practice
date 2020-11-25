@[TOC](无重复字符的最长子串)

## 无重复字符的最长子串

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。


示例 1:
```
输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


```

## 代码
```text

package main

import (
	"fmt"
	"strings"
)

// 无重复字符的最长子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

func main() {
	str := "zxbsadfhjkadfuewjkszdfkljaD"
	n := lengthOfLongestSubstring(str)
	fmt.Println(n)
}

func lengthOfLongestSubstring(s string) int {
	var (
		left   int    = 0
		maxlen int    = 0
		retstr string = ""
	)
	for i, _ := range s {
		if index := strings.Index(retstr, s[i:i+1]); index > -1 {
			left = left + index + 1
		}
		retstr = s[left : i+1]
		if maxlen < len(retstr) {
			maxlen = len(retstr)
		}
	}
	return maxlen
}

```
## 运行结果
```
10
```
## 总结

### 算法

上述的方法最多需要执行 2n 个步骤。事实上，它可以被进一步优化为仅需要 n 个步骤。我们可以定义字符到索引的映射，而不是使用集合来判断一个字符是否存在。 当我们找到重复的字符时，我们可以立即跳过该窗口。

也就是说，如果 s[j] 在 [i, j) 范围内有与 j 重复的字符，我们不需要逐渐增加 i 。
我们可以直接跳过 [i，j'] 范围内的所有元素，并将 i 变为 j + 1。


### 时间复杂度

时间复杂度：O(n)，索引 j 将会迭代 n 次。

空间复杂度（HashMap）：O(min(m,n))，与之前的方法相同。

空间复杂度（Table）：O(m)，m 是字符集的大小。


