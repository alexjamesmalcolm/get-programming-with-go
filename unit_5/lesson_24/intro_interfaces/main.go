package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

func main() {
	m := martian{}
	l := laser(2)
	shout(m)
	shout(l)
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)
}
