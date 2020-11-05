@[TOC](寻找两个有序数组的中位数)

## 寻找两个有序数组的中位数

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。




```
示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

```



## 代码
```text

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	r:=findMedianSortedArrays(nums1,nums2)
	fmt.Println(r)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) > 0 {
		for _, v := range nums2 {
			nums1 = append(nums1, v)
		}
	}
	sort.Ints(nums1)
	length := len(nums1)
	left := length/2 - 1
	if length%2 == 1 {
		return float64(nums1[left+1])
	} else {
		return (float64(nums1[left]) + float64(nums1[left+1])) / 2
	}
}
```


## 运行结果
```
2
```


## 总结

### 算法
为了解决这个问题，我们需要理解 “中位数的作用是什么”。在统计中，中位数被用来：
```
将一个集合划分为两个长度相等的子集，其中一个子集中的元素总是大于另一个子集中的元素。
```

### 时间复杂度

复杂度分析

时间复杂度：O(log(min(m,n)))


空间复杂度：O(1)，
我们只需要恒定的内存来存储 99 个局部变量， 所以空间复杂度为 O(1)。



