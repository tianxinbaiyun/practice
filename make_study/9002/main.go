package main

import (
	"fmt"
)

//定义接口
type adder interface {
	add(string) int
}

//定义函数类型
type handler func(name string) int

//实现函数类型方法
func (h handler) add(name string) int {
	fmt.Println("add:h(name) + 10:", h(name)+10)
	return h(name) + 10
}

//函数参数类型接受实现了adder接口的对象（函数或结构体）
func process(a adder) {
	fmt.Println("process:", a.add("taozs"))
}

//另一个函数定义

func doubler(name string) int {
	fmt.Println("doubler:len(name) * 2:", len(name)*2)
	return len(name) * 2
}

//非函数类型
type myint int

//实现了adder接口
func (i myint) add(name string) int {
	return len(name) + int(i)
}

func main() {
	//注意要成为函数对象必须显式定义handler类型
	//var my handler = func(name string) int {
	//	fmt.Println("my:",len(name))
	//	return len(name)
	//}
	//fmt.Println(my("taozs"))
	fmt.Println(handler(doubler).add("taozs")) //doubler函数显式转换成handler函数对象然后调用对象的add方法
}
