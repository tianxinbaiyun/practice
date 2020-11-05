package shop

type Shop struct {
	User User
}

func NewShop(user User) *Shop {
	return &Shop{
		User: user,
	}
}

func (s Shop) Meeting(guestName string) {
	s.User.Shopping(guestName)
}
