package main

import (
	"fmt"
	"github.com/tianxinbaiyun/practice/try/channel/base"
)

func main() {
	jobs := make(chan int, 1)
	done := make(chan bool)
	fmt.Println(cap(done))
	go func() {
		//      fmt.Println("GoStart")
		for i := 0; ; i++ {
			base.Done(jobs)
		}
	}()
	for j := 1; j <= 10; j++ {
		//      fmt.Println("OutFOR", j)
		base.Add(jobs, j)
		//jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
