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

	SetupPieces(&b)
	return &b
}

func SetupPieces(b *Board) {
	var whiteBishop Piece = &Bishop{Color: WHITE}
	b.SetPiece(b.Sqr("A1"), &whiteBishop)
}

func (b *Board) SetPiece(s *Square, p *Piece) {
	b.Pieces[*s] = *p
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
	col := strings.Index(n, colName)
	if col < 1 || col > 8 {
		fmt.Println("Error, attempting to access col: ", col)
		os.Exit(1)
	}
	return b.GetSquare(row, col+1)
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

func (b *Board) getLine(piece *Piece, startingSquare *Square, rowInc int, colInc int) []Square {

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
		} else if b.isEnemyPieceOn(s, *piece) {
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
	return nil // TODO - implement this
}

func (b *Board) isEnemyPieceOn(square *Square, piece Piece) bool {
	enemyPiece := b.GetPiece(square)
	if enemyPiece != nil {
		return enemyPiece.GetColor() != piece.GetColor()
	}

	return false
}
