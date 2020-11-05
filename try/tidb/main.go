package main

import (
	"github.com/tianxinbaiyun/practice/try/tidb/core/db"
	"github.com/tianxinbaiyun/practice/try/tidb/core/util"
	"github.com/tianxinbaiyun/practice/try/tidb/models"
	"sync"
)

func main() {
	db.Init()
	AddUsers()
}

// AddUsers AddUsers
func AddUsers() {
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				uuid := util.CreateUuid()
				_ = models.Create(&models.Users{
					Name: uuid,
					Age:  j % 100,
				})
			}
		}()
	}
	wg.Wait()
}
