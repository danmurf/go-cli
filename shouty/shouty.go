package shouty

import "strings"

type Shouty interface {

	// Makes a string of text shouty!
	Shout(string) string
}

type loudMouth struct {}

func NewLoudMouth() Shouty {
	return &loudMouth{}
}

func (l loudMouth) Shout(s string) string {
	return strings.ToUpper(s) + "!!!!!"
}
