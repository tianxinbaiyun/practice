package main

import "fmt"

// Stu Stu
type Stu struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*Stu, 0)
	stus := []Stu{
		{Name: "1", Age: 1},
		{Name: "2", Age: 2},
		{Name: "3", Age: 3},
	}

	for _, i2 := range stus {
		tmp := i2
		m[i2.Name] = &tmp
	}

	for _, i2 := range m {
		fmt.Println(i2)
	}
}
