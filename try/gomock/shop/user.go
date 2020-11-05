package shop

import "fmt"

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

func (u *User) Shopping(name string) string {
	return fmt.Sprintf("Hello %s, welcome come to my shop, my shop name is %s", name, u.Name)
}
