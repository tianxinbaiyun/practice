package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	r := findMedianSortedArrays(nums1, nums2)
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
	}
	return (float64(nums1[left]) + float64(nums1[left+1])) / 2
}
