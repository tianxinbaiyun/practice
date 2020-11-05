package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tianxinbaiyun/practice/try/go-micro/api/book/handler"
	"github.com/tianxinbaiyun/practice/try/go-micro/api/book/rpc"
	pb_book "github.com/tianxinbaiyun/practice/try/go-micro/core/pb/book"
	"net/http"
	"strconv"
)

// Book Book
type Book struct {
	handler.Base
}

// List List
func (b *Book) List(c *gin.Context) {

	cateID, _ := strconv.Atoi(c.DefaultQuery("cate_id", "0"))
	subCateID, _ := strconv.Atoi(c.DefaultQuery("sub_cate_id", "0"))
	platform := c.DefaultQuery("platform", "api")

	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "10"))
	curPage, _ := strconv.Atoi(c.DefaultQuery("cur_page", "1"))

	var (
		flag uint32
	)

	rsp, err := rpc.BookSrv.BookList(context.TODO(), &pb_book.Book{
		CateID:    uint32(cateID),
		SubCateID: uint32(subCateID),
		Limit:     uint32(pageNum),
		Offset:    uint32(curPage),
		IsPulish:  1,
		Platform:  platform,
		Flag:      flag,
	})

	if err != nil {
		b.Failure(c, http.StatusInternalServerError, handler.SystemError, "获取数据失败", err)
		return
	}

	data := make([]map[string]interface{}, 0)
	for _, v := range rsp.Data {
		tmp := map[string]interface{}{
			"id":          v.ID,
			"title":       v.Title,
			"cate_id":     v.CateID,
			"sub_cate_id": v.SubCateID,
			"storage":     v.Storage,
			"video":       v.Video,
			"audio":       v.Audio,
			"subject":     v.Subject,
			"book_file":   v.BookFile,
			//"content":     v.Content,
			"is_pulish":  v.IsPulish,
			"hits":       v.Hits,
			"user_id":    v.UserID,
			"author":     v.Author,
			"created_at": v.CreatedAt,
			"updated_at": v.UpdatedAt,
		}
		data = append(data, tmp)
	}

	msg := map[string]interface{}{
		"total":    rsp.Total,
		"page_num": pageNum,
		"cur_page": curPage,
		"data":     data,
	}
	b.Success(c, msg)
}

// Info Info
func (b *Book) Info(c *gin.Context) {
	templeID, _ := strconv.Atoi(c.Param("temple_id"))
	id, _ := strconv.Atoi(c.Param("id"))
	platform := c.DefaultQuery("platform", "api")

	if id == 0 {
		b.Failure(c, http.StatusBadRequest, handler.InvalidParam, "参数id错误", fmt.Errorf("参数错误"))
		return
	}

	rsp, err := rpc.BookSrv.BookInfo(context.TODO(), &pb_book.Book{
		ID:       uint32(id),
		TempleID: uint32(templeID),
		Platform: platform,
	})

	if err != nil {
		b.Failure(c, http.StatusInternalServerError, handler.SystemError, "获取数据失败", err)
		return
	}

	info := rsp.Info
	tmp := map[string]interface{}{
		"id":          info.ID,
		"title":       info.Title,
		"cate_id":     info.CateID,
		"sub_cate_id": info.SubCateID,
		"storage":     info.Storage,
		"video":       info.Video,
		"audio":       info.Audio,
		"subject":     info.Subject,
		"book_file":   info.BookFile,
		"content":     info.Content,
		"is_pulish":   info.IsPulish,
		"hits":        info.Hits,
		"user_id":     info.UserID,
		"author":      info.Author,
		"created_at":  info.CreatedAt,
		"updated_at":  info.UpdatedAt,
	}

	b.Success(c, tmp)
}
