package main

import (
	"testing"
)

func TestCreateBoard(t *testing.T) {
	b := CreateBoard()
	b.Print()

	if b == nil {
		t.Error("Error creating board")
	}
}

func TestSquares(t *testing.T) {
	b := CreateBoard()
	for r := 1; r < 9; r++ {
		for c := 1; c < 9; c++ {
			s := b.GetSquare(r, c)

			if s.Row != r || s.Col != c {
				t.Error("row/col did not match what was expected", r, s.Row)
			}
		}
	}
}

func TestBadSquares(t *testing.T) {
	b := CreateBoard()

	s := b.GetSquare(0, 0)
	if s != nil {
		t.Error("got square with bad index")
	}

	s = b.GetSquare(0, 9)
	if s != nil {
		t.Error("got square with bad index")
	}

	s = b.GetSquare(9, -1)
	if s != nil {
		t.Error("got square with bad index")
	}

	s = b.GetSquare(3, 9)
	if s != nil {
		t.Error("got square with bad index")
	}
}

func TestFourCorners(t *testing.T) {
	b := CreateBoard()
	A1 := b.GetSquare(1, 1)
	A8 := b.GetSquare(8, 1)
	H1 := b.GetSquare(1, 8)
	H8 := b.GetSquare(8, 8)

	if A1.Color != BLACK {
		t.Error("A1 not BLACK")
	}
	if A8.Color != WHITE {
		t.Error("A8 not WHITE")
	}
	if H1.Color != WHITE {
		t.Error("H1 not WHITE")
	}
	if H8.Color != BLACK {
		t.Error("H8 not BLACK")
	}

	if A1.getName() != "A1" {
		t.Error("A1 not named correctly")
	}
	if A8.getName() != "A8" {
		t.Error("A8 not named correctly")
	}
	if H1.getName() != "H1" {
		t.Error("H1 not named correctly")
	}
	if H8.getName() != "H8" {
		t.Error("H8 not named correctly")
	}
}
