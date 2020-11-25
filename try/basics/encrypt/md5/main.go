package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	abc := "123456"
	loop(abc, 10000000)
}

func loop(str string, n int) string {
	var s string
	t := time.Now()
	for i := 1; i < n; i++ {
		h := md5.New()
		h.Write([]byte(str))
		sum := h.Sum(nil)

		//由于是十六进制表示，因此需要转换
		s = hex.EncodeToString(sum)
	}
	fmt.Println("次数", n, "使用时间：", time.Since(t))
	fmt.Println(string(s))
	return s
}
