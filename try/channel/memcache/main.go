package main

import "fmt"

// t t
type T struct {
	i int
}

// produce produce
func produce(ch chan<- T) {
	defer close(ch)
	ch <- T{6}

}

// consume
func consume(ch <-chan T) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan T)
	go produce(ch)
	consume(ch)
}
