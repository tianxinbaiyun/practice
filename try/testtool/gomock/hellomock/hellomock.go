package hellomock

// Talker Talker
type Talker interface {
	SayHello(word string) (response string)
}
