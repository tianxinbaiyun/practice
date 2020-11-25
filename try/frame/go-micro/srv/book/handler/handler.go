package handler

import (
	"context"
	pb_book "github.com/tianxinbaiyun/practice/try/frame/go-micro/core/pb/book"
	"github.com/tianxinbaiyun/practice/try/frame/go-micro/srv/book/db"
)

// BookServiceHandler BookServiceHandler
type BookServiceHandler struct{}

// NewBookService NewBookService
func NewBookService() (h *BookServiceHandler) {
	return &BookServiceHandler{}
}

// BookList BookList
func (b *BookServiceHandler) BookList(ctx context.Context, req *pb_book.Book, rsp *pb_book.BookListRsp) error {
	list, total, err := db.BookList(req)
	if err != nil {
		return err
	}
	rsp.Total = total
	for _, v := range list {
		rsp.Data = append(rsp.Data, v.ToProto())
	}
	return nil
}

// BookInfo BookInfo
func (b *BookServiceHandler) BookInfo(ctx context.Context, req *pb_book.Book, rsp *pb_book.BookInfoRsp) error {
	info, err := db.BooInfo(req)
	if err != nil {
		return err
	}
	rsp.Info = info.ToProto()
	return nil
}

// BookUpdate BookUpdate
func (b *BookServiceHandler) BookUpdate(ctx context.Context, req *pb_book.Book, rsp *pb_book.CommonRsp) error {
	exist, err := db.BookExist(req.ID)
	if err != nil {
		return err
	}
	if !exist {
		rsp.Msg = "数据不存在"
		return nil
	}
	err = db.BookUpdate(req)
	if err != nil {
		return err
	}
	rsp.Ok = true
	return nil
}

// BookStore BookStore
func (b *BookServiceHandler) BookStore(ctx context.Context, req *pb_book.Book, rsp *pb_book.CommonRsp) error {
	err := db.BookStore(req)
	if err != nil {
		return err
	}
	rsp.Ok = true
	return nil
}

// BookDelete BookDelete
func (b *BookServiceHandler) BookDelete(ctx context.Context, req *pb_book.Book, rsp *pb_book.CommonRsp) error {
	exist, err := db.BookExist(req.ID)
	if err != nil {
		return err
	}
	if !exist {
		rsp.Msg = "数据不存在"
		return nil
	}
	err = db.BookDelete(req)
	if err != nil {
		return err
	}
	rsp.Ok = true
	return nil
}
