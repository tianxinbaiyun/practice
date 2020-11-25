package main

import (
	"github.com/tianxinbaiyun/practice/try/mysql/core/db"
	"github.com/tianxinbaiyun/practice/try/mysql/core/util"
	"github.com/tianxinbaiyun/practice/try/mysql/models"
	"sync"
)

func main() {
	db.Init()
	AddUsers()
}

// AddUsers AddUsers
func AddUsers() {
	wg := sync.WaitGroup{}

	for i := 0; i < 900; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				uuid := util.CreateUUID()
				_ = models.Create(&models.Users{
					Name: uuid,
					Age:  j % 100,
				})
			}
		}()
	}
	wg.Wait()
}
