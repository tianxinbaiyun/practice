package hellomock

//Company Company
type Company struct {
	Usher Talker
}

//NewCompany NewCompany
func NewCompany(t Talker) *Company {
	return &Company{
		Usher: t,
	}
}

// Meeting Meeting
func (c Company) Meeting(gusetName string) string {
	return c.Usher.SayHello(gusetName)
}
