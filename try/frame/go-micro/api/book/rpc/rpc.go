package rpc

import (
	"github.com/micro/go-micro/client"
	pb_book "github.com/tianxinbaiyun/practice/try/frame/go-micro/core/pb/book"
)

// 定义
var (
	BookSrv pb_book.BookService
)

func init() {
	BookSrv = pb_book.NewBookService("shanxing.srv.book", client.DefaultClient)
}
