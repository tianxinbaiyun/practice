package main

import (
	"fmt"
	"sync"
	"time"
)

// Record Record
type Record struct {
	ID *int
}

func main() {
	wg := sync.WaitGroup{}
	number := 1000000000
	wg.Add(2)
	go func() {
		defer wg.Done()
		t1 := time.Now()
		for i := 0; i < number; i++ {
			a()
		}
		fmt.Printf("a cost time:%+v\n", time.Since(t1))
	}()
	go func() {
		defer wg.Done()
		t1 := time.Now()
		for i := 0; i < number; i++ {
			b()
		}
		fmt.Printf("b cost time:%+v\n", time.Since(t1))
	}()
	wg.Wait()
}

func a() bool {
	t1 := time.Now()
	return t1.Equal(time.Time{})
}

func b() bool {
	t1 := time.Now()
	return t1.IsZero()
}
