package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	t := time.Now()
	//nums:=[]int{1, 0, -1, 0, -2, 2}
	nums := []int{-3, -2, -1, 0, 0, 0, 0, 1, 2, 3}
	res := fourSum(nums, 0)
	fmt.Println(time.Since(t))
	fmt.Println(res)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var (
		numLen = len(nums)
		left   = numLen - 1
		right  = numLen - 1
		sum    int
		res    [][]int
	)
	res = make([][]int, 0)
	for i := 0; i < left; i++ {
		for j := i + 1; j < right; j++ {
			left = j + 1
			right = numLen - 1
			for left < right {
				sum = nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					obj := []int{nums[i], nums[j], nums[left], nums[right]}
					if !InArray(obj, &res) {
						res = append(res, obj)
					}
					fmt.Printf("i:%v,nums[%v]=%v;", i, i, nums[i])
					fmt.Printf("j:%v,nums[%v]=%v;", j, j, nums[j])
					fmt.Printf("left:%v,nums[%v]=%v;", left, left, nums[left])
					fmt.Printf("right:%v,nums[%v]=%v;\n", right, right, nums[right])

					left++
				}
			}
		}
	}
	return res
}

// InArray InArray
func InArray(obj []int, arr *[][]int) bool {
	for _, v := range *arr {
		fmt.Printf("v:%v,obj:%v;", obj, v)
		if len(obj) == len(v) {
			return true
		}
	}
	return false
}
