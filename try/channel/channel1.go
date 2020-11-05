package main

import (
	"fmt"
	"sync"
	"time"
)

func main1001() {
	var (
		ch1 chan int
		ch2 chan int
	)
	ch1 = make(chan int)
	ch2 = make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			ch1 <- i
			return
		}(i)
	}

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(<-ch1)
		}()
	}
	wg.Wait()
	time.Sleep(time.Hour * 1)
	close(ch1)
	close(ch2)

}

func main1002() {
	jobs := make(chan int, 1)
	done := make(chan bool)
	fmt.Println(cap(done))
	go func() {
		//      fmt.Println("GoStart")
		for i := 1; ; i++ {
			//          fmt.Println("GoforSTART", i)
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
			//fmt.Println("GoforEND", i)
		}
	}()
	for j := 1; j <= 3; j++ {
		//      fmt.Println("OutFOR", j)
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
