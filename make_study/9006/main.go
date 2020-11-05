package main

import "fmt"

func main() {
	// 以两倍扩容
	s := []int{1, 2, 3, 4}
	printSlice(s)

	s = append(s, 9)
	printSlice(s)

	s = append(s, 11)
	printSlice(s)
	// 2倍扩容也不满足时
	s1 := []int{1}
	s1 = append(s1, 1, 2, 3)
	printSlice(s1)

	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	printSlice(s1)
}

func printSlice(s []int) {
	fmt.Printf("slice:%+v,point:%p,len:%d,cap:%d\n", s, &s, len(s), cap(s))
}
