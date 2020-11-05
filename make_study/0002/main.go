package main

import "os"
import "fmt"

// Human Human
type Human interface {
	sayHello()
}

//Chinese Chinese
type Chinese struct {
	name string
}

// English English
type English struct {
	name string
}

// sayHello sayHello
func (c *Chinese) sayHello() {
	fmt.Println(c.name, "说：你好，世界")
}

func (e *English) sayHello() {
	fmt.Println(e.name, "says: hello,world")
}

func main() {
	fmt.Println(len(os.Args))

	c := Chinese{"汪星人"}
	e := English{"jorn"}
	m := map[int]Human{}

	m[0] = &c
	m[1] = &e

	for i := 0; i < 2; i++ {
		m[i].sayHello()
	}
}
