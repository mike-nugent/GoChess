package main

const (
	WP = "♟"
	WR = "♜"
	WN = "♞"
	WB = "♝"
	WQ = "♛"
	WK = "♚"

	BP = "♙"
	BR = "♖"
	BN = "♘"
	BB = "♗"
	BQ = "♕"
	BK = "♔"
)

const (
	WHITE = "white"
	BLACK = "black"
)

const (
	KING   = iota
	QUEEN  = iota
	ROOK   = iota
	BISHOP = iota
	KNIGHT = iota
	PAWN   = iota
)

type Piece interface {
	PieceType() int
	ValidMoves(startSquare string) []string
	GetColor() string
}

func GetPiece(str string, color string) int {
	if color == WHITE {
		if str == WP {
			return PAWN
		} else if str == WK {
			return KING
		} else if str == WR {
			return ROOK
		} else if str == WB {
			return BISHOP
		} else if str == WN {
			return KNIGHT
		} else if str == WQ {
			return QUEEN
		}
	} else {
		if str == BP {
			return PAWN
		} else if str == BK {
			return KING
		} else if str == BR {
			return ROOK
		} else if str == BB {
			return BISHOP
		} else if str == BN {
			return KNIGHT
		} else if str == BQ {
			return QUEEN
		}
	}

	return -1
}
