package models

import (
	"github.com/tianxinbaiyun/practice/try/tidb/core/db"
	"github.com/tianxinbaiyun/practice/try/tidb/core/slog"
)

// Users Users
type Users struct {
	UserID    int    `json:"user_id" xorm:"not null pk autoincr INT(11)"`
	Name      string `json:"name" xorm:"not null default '' index VARCHAR(200)"`
	Age       int    `json:"age" xorm:"not null default 0 INT(11)"`
	CreatedAt int    `json:"created_at" xorm:"created not null default 0 INT(11)"`
	UpdatedAt int    `json:"updated_at" xorm:"updated not null default 0 INT(11)"`
}

// Page Page
type Page struct {
	Limit  uint32 `json:"limit,omitempty"`
	Offset uint32 `json:"offset,omitempty"`
	IsPage bool   `json:"is_page,omitempty"`
}

// UserListReq UserListReq
type UserListReq struct {
	Users *Users `json:"users"`
	Page  *Page  `json:"page"`
}

// Create Create
func Create(par *Users) (err error) {
	q := db.Get().Table("users")
	_, err = q.Insert(par)
	return
}

// Update Update
func Update(par *Users) (err error) {
	q := db.Get().Table("users")
	q.Where("user_id = ?", par.UserID)
	_, err = q.Update(par)
	return
}

// Delete Delete
func Delete(par *Users) (err error) {
	q := db.Get().Table("users")
	q.Where("user_id = ?", par.UserID)
	_, err = q.Delete(&Users{})
	return
}

// GetList GetList
func GetList(par *UserListReq) (list []*Users, total int64, err error) {
	list = make([]*Users, 0)
	q := db.Get().Table("users")
	if par.Users != nil {
		if par.Users.UserID > 0 {
			q.Where("user_id =?", par.Users.UserID)
		}
		if par.Users.Name != "" {
			q.Where("name like ?", "%"+par.Users.Name+"%")
		}
	}
	total, err = q.Clone().Count(&Users{})
	if err != nil {
		slog.Error(err)
		return
	}
	if par.Page != nil {
		if par.Page.Limit == 0 {
			par.Page.Limit = 10
		}
		if par.Page.Offset == 0 {
			par.Page.Offset = 1
		}
		limit := int(par.Page.Limit)
		offset := (int(par.Page.Offset) - 1) * limit
		q.Limit(limit, offset)
	}

	err = q.OrderBy("user_id asc").Find(&list)
	if err != nil {
		slog.Error(err)
	}

	return
}
