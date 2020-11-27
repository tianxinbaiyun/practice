package main

import (
	"fmt"
	"log"
)

type A struct {
	int
}

func main() {
	var (
		dataInt       int
		dataFloat     float64
		dataComplex64 complex64
		dataBool      bool
		dataByte      byte
		dataString    string
		dataError     error
		dataMap       map[int]int
		dataSlice     []int
		dataArray     [1]int
		dataChan      chan int
		dataInterface interface{}
		dataPoint     *int
		dataStruct    A
		dataFunc      func()
	)

	dataInt = 1
	log.Println(dataInt)
	dataFloat = 1
	log.Println(dataFloat)
	dataComplex64 = 1
	log.Println(dataComplex64)
	dataBool = true
	log.Println(dataBool)
	dataByte = 1
	log.Println(dataByte)
	dataString = "1"
	log.Println(dataString)
	dataError = fmt.Errorf("1")
	log.Println(dataError)
	dataMap = map[int]int{1: 1}
	log.Println(dataMap)
	dataSlice = []int{1}
	log.Println(dataSlice)
	dataArray = [1]int{2}
	log.Println(dataArray)
	dataChan = make(chan int, 10)
	dataChan <- 1
	log.Println(<-dataChan)
	dataInterface = 1
	log.Println(dataInterface)
	dataPoint = &dataInt
	log.Println(dataPoint)
	dataStruct = A{1}
	log.Println(dataStruct)
	dataFunc = func() {
		fmt.Println(111)
	}
	log.Println(dataFunc)
}
