package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unsafe"
)

func main() {
	//var(
	//	str string = ""
	//	n int = 100000
	//)
	t := time.Now()
	test2()
	fmt.Println(time.Since(t))
	time.Sleep(time.Hour * 1)
}

func test1() {
	userFile := "/home/clx/txt/1.txt"
	//读文件
	readbuf, _ := ioutil.ReadFile(userFile)
	//文字数据处理
	strcontent := string(readbuf)
	for i := 0; i < 100; i++ {
		strcontent = ConnectString1(strcontent, strcontent+"和气生财")
	}
	buf := []byte(strcontent)
	//写文件
	ioutil.WriteFile(userFile, buf, os.ModeExclusive)
	fmt.Println("end")
}

func test2() {
	userFile := "/home/clx/txt/1.txt"
	//读文件
	readbuf, _ := ioutil.ReadFile(userFile)
	//文字数据处理
	strcontent := string(readbuf)
	for i := 0; i < 5; i++ {

		strcontent = ConnectString2(strcontent, strcontent+"和气生财")
	}
	buf := []byte(strcontent)
	//写文件
	ioutil.WriteFile(userFile, buf, os.ModeExclusive)
}

// ConnectString ConnectString
func ConnectString(str1, str2 *string) (str *string) {
	*str = ConnectString2(*str1, *str2)
	return
}

//ConnectString1 使用+连接
func ConnectString1(str1, str2 string) string {
	return str1 + str2
}

// ConnectString2 使用Builder连接
func ConnectString2(str1, str2 string) string {
	var sb strings.Builder
	sb.WriteString(str1)
	sb.WriteString(str2)
	return sb.String()
}

// Str2bytes 字符串转字符数组
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2str 字符数组转字符串
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
