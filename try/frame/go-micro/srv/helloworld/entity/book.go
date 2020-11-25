package entity

import pb_book "github.com/tianxinbaiyun/practice/try/frame/go-micro/core/pb/book"

// Book Book
type Book struct {
	ID         uint32 `json:"id" xorm:"not null pk autoincr INT(10)"`
	Title      string `json:"title"`
	CateID     uint32 `json:"cate_id"`
	SubCateID  uint32 `json:"sub_cate_id"`
	Storage    uint32 `json:"storage"`
	Video      string `json:"video"`
	Audio      uint32 `json:"audio"`
	BookFile   uint32 `json:"book_file"`
	Subject    string `json:"subject"`
	Content    string `json:"content"`
	IsPulish   int32  `json:"is_pulish"`
	Hits       uint32 `json:"hits"`
	TempleID   uint32 `json:"temple_id"`
	UserID     uint32 `json:"user_id"`
	Flag       uint8  `json:"flag"`
	Author     string `json:"author"`
	CreatedAt  string `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt  string `json:"updated_at" xorm:"updated TIMESTAMP"`
	HasCollect bool   `json:"has_collect" xorm:"-"`
}

// ToProto ToProto
func (b *Book) ToProto() *pb_book.Book {
	return &pb_book.Book{
		ID:         b.ID,
		Title:      b.Title,
		CateID:     b.CateID,
		SubCateID:  b.SubCateID,
		Storage:    b.Storage,
		Video:      b.Video,
		Audio:      b.Audio,
		BookFile:   b.BookFile,
		Subject:    b.Subject,
		Content:    b.Content,
		IsPulish:   b.IsPulish,
		Hits:       b.Hits,
		TempleID:   b.TempleID,
		UserID:     b.UserID,
		Author:     b.Author,
		CreatedAt:  b.CreatedAt,
		UpdatedAt:  b.UpdatedAt,
		HasCollect: b.HasCollect,
		Platform:   "",
		Flag:       0,
	}
}
