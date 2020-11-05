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
		res     int = 0
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

func threeSumClosest2(nums []int, target int) int {
	sort.Ints(nums)
	//fmt.Println(nums)
	var (
		numsLen int = len(nums)
		left    int = 0
		right   int = numsLen - 1
		res     int = 0
	)
	res = nums[0] + nums[1] + nums[2]

	for i := 0; i < numsLen-2; i++ {
		left = i + 1
		right = numsLen - 1
		for left != right {
			min := nums[i] + nums[left] + nums[left+1]
			if target < min {
				if math.Abs(float64(min-target)) < math.Abs(float64(res-target)) {
					res = min
				}
				break
			}
			max := nums[i] + nums[right-1] + nums[right]
			if max < target {
				if math.Abs(float64(max-target)) < math.Abs(float64(res-target)) {
					res = max
				}
				break
			}
			sum := nums[i] + nums[left] + nums[right]
			fmt.Println(sum)
			if sum == target {
				return sum
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
			}
			if sum < target {
				left++
				for left != right && nums[left] == nums[left+1] {
					left++
				}
			} else {
				right--
				for left != right && nums[right] == nums[right+1] {
					right--
				}
			}
		}

	}

	return res
}
