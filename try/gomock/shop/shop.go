package shop

// Shop Shop
type Shop struct {
	User User
}

// NewShop NewShop
func NewShop(user User) *Shop {
	return &Shop{
		User: user,
	}
}

// Meeting Meeting
func (s Shop) Meeting(guestName string) {
	s.User.Shopping(guestName)
}
