package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type Stu struct {
	Name string
	Age  int
}

func main() {
	pprof()
	test()
}
func pprof() {
	go func() {
		// 开启pprof，监听请求
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
		}
	}()
}

func test() {
	m := make(map[string]*Stu, 0)
	stus := []Stu{
		{Name: "1", Age: 1},
		{Name: "2", Age: 2},
		{Name: "3", Age: 3},
	}
	lock := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			for _, i2 := range stus {
				tmp := i2
				m[i2.Name] = &tmp
			}

			for _, i2 := range m {
				fmt.Println(i2)
			}
		}()
	}
	time.Sleep(time.Hour * 1)
}
