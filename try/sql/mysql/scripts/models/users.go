package models

//Users Users
type Users struct {
	UserID    int    `json:"user_id" xorm:"not null pk autoincr INT(11)"`
	Name      string `json:"name" xorm:"not null default '' index VARCHAR(200)"`
	Age       int    `json:"age" xorm:"not null default 0 INT(11)"`
	CreatedAt int    `json:"created_at" xorm:"not null default 0 INT(11)"`
	UpdatedAt int    `json:"updated_at" xorm:"not null default 0 INT(11)"`
}
