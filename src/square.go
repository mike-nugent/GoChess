package main

import (
	"strconv"
)

type Square struct {
	Color string
	Row   int
	Col   int
}

func (s *Square) toString() string {
	return s.getName() + " " + strconv.Itoa(s.Row) + " " + strconv.Itoa(s.Col) + " " + s.Color
}

func (s *Square) getName() string {
	letters := "ABCDEFGH"
	return string(letters[s.Col-1]) + strconv.Itoa(s.Row)
}
