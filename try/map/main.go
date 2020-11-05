package main

import "fmt"

func main() {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	AddMap(m)
	fmt.Println(m)
	Remove(m)
	fmt.Println(m)
}

//AddMap AddMap
func AddMap(m map[int]int) {
	m[4] = 123
}

// Remove Remove
func Remove(m map[int]int) {
	delete(m, 2)
}
