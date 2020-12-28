package main

import (
	"fmt"
	"strings"
)

func (game *Game) ValidateMove(move string) (bool, string) {

	fmt.Println("move: " + move)
	pieceToMoveType := -1
	squareToMoveTo := ""
	//	pieceToMove := ""
	//	currentSquare := ""
	//	newSquare := ""

	if move == "0-0" || move == "0-0-0" {
		fmt.Println("TODO - handle castling")
		return false, "castling not implemented yet"
	} else if strings.Contains(move, "x") {
		fmt.Println("TODO - implement capture")
		return false, "capturing not implemented yet"
	} else if strings.Contains(move, "=") {
		fmt.Println("TODO - implement promotion")
		return false, "promotion not implemented yet"
	} else {
		// Most likely it's just a piece move.  Break it down to the piece and the destination
		pieceToMoveType = getPieceType(move)
		squareToMoveTo = getDestinationSquare(move)
		squaresToMoveFrom, errMsg := getOriginationSquare(move, squareToMoveTo, game.Board, game.CurrentTurn)
		if errMsg != "" {
			return false, errMsg + fmt.Sprint(squaresToMoveFrom)
		}
		fmt.Println(squareToMoveTo)
		fmt.Println(pieceToMoveType)
		fmt.Println(squaresToMoveFrom, errMsg)

	}
	//	pieceToMove := getPiece(move)

	//square := ""

	//TODO add corrected move back, in case of king in check +

	return true, "dunno"
}

func getOriginationSquare(move string, destinationSquare string, board map[string]string, color string) ([]string, string) {
	direction := 1
	if color == BLACK {
		direction = -1
	}

	pieceType := getPieceType(move)
	squares := getSquaresOfType(pieceType, color, board)

	for square, _ := range squares {
		var matches []string

		if CanMoveTo(pieceType, square, destinationSquare) {
			matches = append(matches, square)
		}
		errorMessage := ""
		if len(matches) > 1 {
			//More than one piece can move to this square. In this case, the move is ambiguous and needs addition row/col information in the notation
			errorMessage = "Ambiguous move. More than 1 piece can move here. Provide more specific notation."
		} else if len(matches) == 0 {
			//No pieces can move here.  In this case, this is an invalid move and should be reported as such.
			errorMessage = "No piece can move here.  Check notation."

		}

		return matches, errorMessage
	}

	fmt.Println(squares, direction)
	return nil, "error - unhandled"
}

func CanMoveTo(pieceType int, fromSquare string, toSquare string) bool {
	return true // TODO fix this
}

func getSquaresOfType(piece int, color string, board map[string]string) map[string]int {
	squares := make(map[string]int)
	for key, val := range board {
		if GetPiece(val, color) == piece {
			squares[key] = piece
		}
	}
	return squares
}

func getDestinationSquare(move string) string {
	if strings.Contains(move, "x") {
		pre := move[strings.Index(move, "x")+1:]
		if strings.HasSuffix(pre, "+") || strings.HasSuffix(pre, "#") {
			return pre[:len(pre)-1]
		}
		return pre
	} else {
		return move[len(move)-2:]
	}
}

func getPieceType(move string) int {

	fc := move[:1]

	if fc == "Q" {
		return QUEEN
	} else if fc == "K" {
		return KING
	} else if fc == "R" {
		return ROOK
	} else if fc == "B" {
		return BISHOP
	} else if fc == "N" {
		return KNIGHT
	} else {
		return PAWN
	}
}
