package handler

import (
	"context"
	pb_book "github.com/tianxinbaiyun/practice/try/go-micro/core/pb/book"
	"github.com/tianxinbaiyun/practice/try/go-micro/srv/book/db"
)

type bookServiceHandler struct{}

func NewBookService() *bookServiceHandler {
	return &bookServiceHandler{}
}

func (b *bookServiceHandler) BookList(ctx context.Context, req *pb_book.Book, rsp *pb_book.BookListRsp) error {
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

func (b *bookServiceHandler) BookInfo(ctx context.Context, req *pb_book.Book, rsp *pb_book.BookInfoRsp) error {
	info, err := db.BooInfo(req)
	if err != nil {
		return err
	}
	rsp.Info = info.ToProto()
	return nil
}

func (b *bookServiceHandler) BookUpdate(ctx context.Context, req *pb_book.Book, rsp *pb_book.CommonRsp) error {
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

func (b *bookServiceHandler) BookStore(ctx context.Context, req *pb_book.Book, rsp *pb_book.CommonRsp) error {
	err := db.BookStore(req)
	if err != nil {
		return err
	}
	rsp.Ok = true
	return nil
}

func (b *bookServiceHandler) BookDelete(ctx context.Context, req *pb_book.Book, rsp *pb_book.CommonRsp) error {
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
