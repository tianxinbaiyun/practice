package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println(time.Now())
		time.Sleep(30 * time.Second)
	}
}
