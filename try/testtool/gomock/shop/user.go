package shop

import "fmt"

// User User
type User struct {
	Name string
}

// NewUser NewUser
func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

// Shopping Shopping
func (u *User) Shopping(name string) string {
	return fmt.Sprintf("Hello %s, welcome come to my shop, my shop name is %s", name, u.Name)
}
