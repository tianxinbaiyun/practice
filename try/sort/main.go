package main

import "fmt"

func main() {
	var aa *string
	bb := "25"
	aa = &bb
	fmt.Println(*aa)
}
