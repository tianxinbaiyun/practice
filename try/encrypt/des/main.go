package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"github.com/tianxinbaiyun/practice/try/mysql/core/util"
	"time"
)

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func encrypt3DES(src []byte, key []byte) []byte {
	block, _ := des.NewTripleDESCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	blockmode.CryptBlocks(src, src)
	return src
}

func decrypt3DES(src []byte, key []byte) []byte {
	block, _ := des.NewTripleDESCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}

func main() {
	//	x := []byte("你中了女神徐莉的毒，唯一解锁的方式是和她在一起。")
	//	key := []byte("87654321hgfedcbaopqrstuv")
	//	x1 := encrypt3DES(x, key)
	//	x2 := decrypt3DES(x1, key)
	//	fmt.Print(string(x2))
	str := "123456"
	loop(str, 10000000)
}

const key = "87654321hgfedcbaopqrstuv"

func loop(str string, n int) string {
	var s string
	var b []byte
	t := time.Now()
	for i := 1; i < n; i++ {
		b = encrypt3DES(util.Str2bytes(str), util.Str2bytes(key))
	}
	fmt.Println("次数", n, "使用时间：", time.Since(t))
	fmt.Println(string(b))
	return s
}
