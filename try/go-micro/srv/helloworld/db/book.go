package db

import (
	"github.com/tianxinbaiyun/practice/try/go-micro/core/db"
	pb_book "github.com/tianxinbaiyun/practice/try/go-micro/core/pb/book"
	"github.com/tianxinbaiyun/practice/try/go-micro/core/slog"
	"github.com/tianxinbaiyun/practice/try/go-micro/srv/book/entity"
)

// BookList BookList
func BookList(par *pb_book.Book) (list []*entity.Book, total int64, err error) {
	list = make([]*entity.Book, 0)
	q := db.Get().Table("book").Omit("content")
	if par.CateID > 0 {
		q.Where("cate_id = ?", par.CateID)
	}
	if par.SubCateID > 0 {
		q.Where("sub_cate_id = ?", par.SubCateID)
	}

	if par.Title != "" {
		q.Where("title like ?", "%"+par.Title+"%")
	}

	if par.IsPulish != -1 { //-1则不过滤
		q.Where("is_pulish = ?", par.IsPulish)
	}

	total, err = q.Clone().Count(entity.Book{})
	if err != nil {
		slog.Error(err)
		return
	}
	if par.Limit == 0 {
		par.Limit = 10
	}
	if par.Offset == 0 {
		par.Offset = 1
	}

	limit := int(par.Limit)
	offset := (int(par.Offset) - 1) * limit
	if par.CateID == 2 {
		q.OrderBy("created_at asc")
	} else {
		q.OrderBy("created_at desc")
	}

	err = q.Limit(limit, offset).Find(&list)
	if err != nil {
		slog.Error(err)
	}
	return
}

// BooInfo BooInfo
func BooInfo(par *pb_book.Book) (info *entity.Book, err error) {
	info = new(entity.Book)
	q := db.Get().Table("book")
	q.Where("id = ?", par.ID)
	_, err = q.Get(info)
	if err != nil {
		slog.Error(err)
	}
	return
}

type masterName struct {
	Name string
}

// BookStore BookStore
func BookStore(par *pb_book.Book) error {
	//获取法师的名称
	master := new(masterName)
	_, _ = db.Get().Table("temple_masters").Select("name").
		Where("user_id = ? and temple_id = ?", par.UserID, par.TempleID).Get(master)
	if par.Platform == "sxadmin" {
		par.TempleID = 0
	}

	var (
		flag uint8
	)
	if par.CateID == 2 {
		if par.Content != "" && par.BookFile == 0 { //经书内容不为空且经书文件id为空
			flag = 1
		}
	}
	insertData := &entity.Book{
		Title:     par.Title,
		CateID:    par.CateID,
		SubCateID: par.SubCateID,
		Storage:   par.Storage,
		Video:     par.Video,
		Audio:     par.Audio,
		BookFile:  par.BookFile,
		Subject:   par.Subject,
		Content:   par.Content,
		IsPulish:  par.IsPulish,
		TempleID:  par.TempleID,
		UserID:    par.UserID,
		Author:    master.Name,
		Flag:      flag,
	}
	q := db.Get().Table("book")
	_, err := q.Insert(insertData)
	if err != nil {
		slog.Error(err)
	}
	return err
}

// BookUpdate BookUpdate
func BookUpdate(par *pb_book.Book) error {
	updateData := &entity.Book{
		Title:     par.Title,
		CateID:    par.CateID,
		SubCateID: par.SubCateID,
		Storage:   par.Storage,
		Video:     par.Video,
		Audio:     par.Audio,
		BookFile:  par.BookFile,
		Subject:   par.Subject,
		Content:   par.Content,
		IsPulish:  par.IsPulish,
	}

	if par.CateID == 2 {
		if par.Content != "" && par.BookFile == 0 { //经书内容不为空且经书文件id为空
			updateData.Flag = 1
		}
	}
	q := db.Get().Table("book")
	q.Where("id = ?", par.ID)
	q.Where("temple_id = ?", par.TempleID)

	_, err := q.
		Cols("title", "cate_id", "sub_cate_id", "storage", "video", "audio", "book_file", "subject", "content", "is_pulish", "flag").
		Update(updateData)
	if err != nil {
		slog.Error(err)
	}
	return err
}

// BookDelete BookDelete
func BookDelete(par *pb_book.Book) error {
	q := db.Get().Table("book")
	q.Where("id = ?", par.ID)
	_, err := q.Where("temple_id = ?", par.TempleID).Delete(&entity.Book{})
	if err != nil {
		slog.Error(err)
	}
	return err
}

// BookExist BookExist
func BookExist(id uint32) (exist bool, err error) {
	exist, err = db.Get().Table("book").Where("id = ?", id).Exist(&entity.Book{})
	if err != nil {
		slog.Error(err)
	}
	return
}
