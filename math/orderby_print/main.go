package main

import (
	"fmt"
	"sync"
)

// PrintTest PrintTest
type PrintTest struct {
}

func (PrintTest) print1() {
	fmt.Println(1)
}
func (PrintTest) print2() {
	fmt.Println(2)
}

func (PrintTest) print3() {
	fmt.Println(3)
}

func main() {
	var (
		wg1 sync.WaitGroup
		wg2 sync.WaitGroup
		wg3 sync.WaitGroup
		t   PrintTest
	)
	wg1.Add(1)
	wg2.Add(1)
	wg3.Add(1)
	fmt.Println("fist:", wg1, wg2, wg3)
	go func() {
		fmt.Println(111)
		t.print1()
		fmt.Println(wg1)
		wg1.Done()
	}()
	go func() {
		wg1.Wait()
		fmt.Println(222)
		t.print2()
		wg2.Done()
		fmt.Println(wg2)
	}()
	go func() {
		wg2.Wait()
		fmt.Println(333)
		t.print3()
		wg3.Done()
		fmt.Println(wg3)
	}()
	wg3.Wait()
}
