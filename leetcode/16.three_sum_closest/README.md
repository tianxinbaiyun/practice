@[TOC](最接近的三数之和)

## 最接近的三数之和

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。

返回这三个数的和。假定每组输入只存在唯一答案。


```
例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).


```



## 代码
```text
package main

import (
	"fmt"
	"math"
	"sort"
)



func main() {
	//input := []int{1, 2, 4, 6, 8, 110, 12, 43456, 1111, 22, 23}
	//input := []int{0, 1, 2}
	//input := []int{0, -1}
	//input := []int{1, 2, 4, 8, 16, 32, 64, 128}
	//input := []int{1, 1, 1, 1}
	input := []int{6, -18, -20, -7, -15, 9, 18, 10, 1, -20, -17, -19, -3, -5, -19, 10, 6, -11, 1, -17, -15, 6, 17, -18, -3, 16, 19, -20, -3, -17, -15, -3, 12, 1, -9, 4, 1, 12, -2, 14, 4, -4, 19, -20, 6, 0, -19, 18, 14, 1, -15, -5, 14, 12, -4, 0, -10, 6, 6, -6, 20, -8, -6, 5, 0, 3, 10, 7, -2, 17, 20, 12, 19, -13, -1, 10, -1, 14, 0, 7, -3, 10, 14, 14, 11, 0, -4, -15, -8, 3, 2, -5, 9, 10, 16, -4, -3, -9, -8, -14, 10, 6, 2, -12, -7, -16, -6, 10}
	output := threeSumClosest(input, -52)
	fmt.Println(output)
	return
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		res     int
		numsLen int = len(nums)
	)
	res = nums[0] + nums[1] + nums[2]
	for i := 0; i < numsLen-2; i++ {
		left := i + 1
		right := numsLen - 1
		for left != right {
			min := nums[i] + nums[left] + nums[left+1]
			if target < min {
				if math.Abs(float64(res-target)) > math.Abs(float64(min-target)) {
					res = min
				}
				break
			}
			max := nums[i] + nums[right] + nums[right-1]
			if target > max {
				if math.Abs(float64(res-target)) > math.Abs(float64(max-target)) {
					res = max
				}
				break
			}
			sum := nums[i] + nums[left] + nums[right]
			//fmt.Println(sum)
			if sum == target {
				return sum
			}
			if math.Abs(float64(res-target)) > math.Abs(float64(sum-target)) {
				res = sum
			}
			if sum < target {
				left++
				for left != right && nums[left] == nums[left+1] {
					left++
				}
			} else {
				right--
				for left != right && nums[right] == nums[right-1] {
					right--
				}
			}
		}
	}
	return res
}

```


## 运行结果

```
-52
```


## 总结

### 算法

排序和双指针

在数组 nums 中，进行遍历，每遍历一个值利用其下标i，形成一个固定值 nums[i]

再使用前指针指向 start = i + 1 处，后指针指向 end = nums.length - 1 处，也就是结尾处

根据 sum = nums[i] + nums[start] + nums[end] 的结果，判断 sum 与目标 target 的距离，如果更近则更新结果 ans

同时判断 sum 与 target 的大小关系，因为数组有序，如果 sum > target 则 end--，如果 sum < target 则 start++，如果 sum == target 则说明距离为 0 直接返回结果



### 时间复杂度

本题目因为要计算三个数，如果靠暴力枚举的话时间复杂度会到 O(n^3)，需要降低时间复杂度
首先进行数组排序，时间复杂度 O(nlogn)

整个遍历过程，固定值为 n 次，双指针为 n 次，时间复杂度为 O(n^2)
总时间复杂度：O(nlogn) + O(n^2) = O(n^2)













