package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

// paseStudent paseStudent
func paseStudent() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		fmt.Printf("%v\t%p\n", stu, &stu)
		tmp := stu
		m[stu.Name] = &tmp
	}
	return m
}
func main() {
	students := paseStudent()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
