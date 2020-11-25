package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	slice = append(slice, 1)
	slice = appendSlide(slice)
	fmt.Printf("slide,len:%d,cap:%d,point:%p", len(slice), cap(slice), &slice)
	fmt.Println(slice)
}

func appendSlide(s []int) []int {
	s[1] = 12
	fmt.Printf("s,len:%d,cap:%d,point:%p", len(s), cap(s), &s)
	s = append(s, 1)
	fmt.Printf("s,len:%d,cap:%d,point:%p", len(s), cap(s), &s)
	fmt.Println(s)
	return s
}
