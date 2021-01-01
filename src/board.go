package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Squares [8][8]Square
	Pieces  map[Square]Piece
}

func CreateBoard() *Board {
	var b Board
	color := BLACK

	for r := 0; r < 8; r++ {
		if color == WHITE {
			color = BLACK
		} else {
			color = WHITE
		}
		for c := 0; c < 8; c++ {
			if color == WHITE {
				color = BLACK
			} else {
				color = WHITE
			}
			b.Squares[r][c] = Square{Color: color, Row: r + 1, Col: c + 1}
		}
	}

	b.Pieces = make(map[Square]Piece)
	return &b
}

func (b *Board) NewGame() {
	var whiteBishop Piece = &Bishop{Color: WHITE}
	var blackBishop Piece = &Bishop{Color: BLACK}
	// White
	b.SetPiece(b.Sqr("A1"), nil)
	b.SetPiece(b.Sqr("B1"), nil)
	b.SetPiece(b.Sqr("C1"), &whiteBishop)
	b.SetPiece(b.Sqr("D1"), nil)
	b.SetPiece(b.Sqr("E1"), nil)
	b.SetPiece(b.Sqr("F1"), &whiteBishop)
	b.SetPiece(b.Sqr("G1"), nil)
	b.SetPiece(b.Sqr("H1"), nil)

	b.SetPiece(b.Sqr("A2"), nil)
	b.SetPiece(b.Sqr("B2"), nil)
	b.SetPiece(b.Sqr("C2"), nil)
	b.SetPiece(b.Sqr("D2"), nil)
	b.SetPiece(b.Sqr("E2"), nil)
	b.SetPiece(b.Sqr("F2"), nil)
	b.SetPiece(b.Sqr("G2"), nil)
	b.SetPiece(b.Sqr("H2"), nil)

	//Black
	b.SetPiece(b.Sqr("A8"), nil)
	b.SetPiece(b.Sqr("B8"), nil)
	b.SetPiece(b.Sqr("C8"), &blackBishop)
	b.SetPiece(b.Sqr("D8"), nil)
	b.SetPiece(b.Sqr("E8"), nil)
	b.SetPiece(b.Sqr("F8"), &blackBishop)
	b.SetPiece(b.Sqr("G8"), nil)
	b.SetPiece(b.Sqr("H8"), nil)

	b.SetPiece(b.Sqr("A7"), nil)
	b.SetPiece(b.Sqr("B7"), nil)
	b.SetPiece(b.Sqr("C7"), nil)
	b.SetPiece(b.Sqr("D7"), nil)
	b.SetPiece(b.Sqr("E7"), nil)
	b.SetPiece(b.Sqr("F7"), nil)
	b.SetPiece(b.Sqr("G7"), nil)
	b.SetPiece(b.Sqr("H7"), nil)
}

func (b *Board) SetPiece(s *Square, p *Piece) {
	if p != nil {
		b.Pieces[*s] = *p
	}
}

func (b *Board) Sqr(name string) *Square {
	if len(name) != 2 {
		fmt.Println("Error, attempting to access square: " + name)
		os.Exit(1)
	}
	row, _ := strconv.Atoi(string(name[1]))
	colName := string(name[0])
	if row < 1 || row > 8 {
		fmt.Println("Error, attempting to access row: ", row)
		os.Exit(1)
	}
	n := "ABCDEFGH"
	col := strings.Index(n, colName) + 1
	if col < 1 || col > 8 {
		fmt.Println("Error, attempting to access col: ", col)
		os.Exit(1)
	}
	return b.GetSquare(row, col)
}

func (b *Board) GetSquare(r int, c int) *Square {

	r -= 1
	c -= 1
	if r < 0 || r > 7 || c < 0 || c > 7 {
		return nil
	}

	return &b.Squares[r][c]
}

func (b *Board) Print() {
	for r := 1; r < 9; r++ {
		for c := 1; c < 9; c++ {
			fmt.Println(b.GetSquare(r, c).toString())
		}
	}
}

func (b *Board) getLine(color string, startingSquare *Square, rowInc int, colInc int) []Square {

	var sqrs []Square
	r := 0
	c := 0
	for i := 0; i < 8; i++ {
		r += rowInc
		c += colInc

		s := b.GetSquare(startingSquare.Row+r, startingSquare.Col+c)
		if s == nil {
			return sqrs
		} else if b.isEmpty(s) {
			sqrs = append(sqrs, *s)
		} else if b.isEnemyPieceOn(s, color) {
			sqrs = append(sqrs, *s)
			return sqrs
		} else {
			return sqrs
		}
	}
	return sqrs
}

func (b *Board) isEmpty(square *Square) bool {

	return b.GetPiece(square) == nil
}

func (b *Board) GetPiece(square *Square) Piece {
	return b.Pieces[*square]
}

func (b *Board) isEnemyPieceOn(square *Square, color string) bool {
	enemyPiece := b.GetPiece(square)
	if enemyPiece != nil {
		return enemyPiece.GetColor() != color
	}

	return false
}
