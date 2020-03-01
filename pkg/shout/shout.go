package shout

import "strings"

type Shout interface {
	Shout(string, int) string
}

type shout struct{}

func NewShout() Shout {
	return &shout{}
}

func (l shout) Shout(s string, lv int) string {
	return strings.ToUpper(s) + strings.Repeat("!", lv)
}
