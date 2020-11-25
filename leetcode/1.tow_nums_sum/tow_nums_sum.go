package main

//两数之和
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	var tmp map[int]int
	tmp = make(map[int]int, 0)
	for i, i2 := range nums {
		if _, ok := tmp[target-i2]; ok {
			return []int{tmp[target-i2], i}
		}
		tmp[i2] = i
	}
	return nil
}

func twoSum0(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
