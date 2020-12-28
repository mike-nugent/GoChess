package main

import (
	"strconv"
	"testing"
)

func TestGetDestinationSquare(t *testing.T) {
	Check(t, "e4", getDestinationSquare("e4"), "e4")
	Check(t, "d4", getDestinationSquare("Nd4"), "Nd4")
	Check(t, "h1", getDestinationSquare("dxh1"), "dxh1")
	Check(t, "a3", getDestinationSquare("Kxa3"), "Kxa3")
	Check(t, "f2", getDestinationSquare("Qf2"), "Qf2")
	Check(t, "b5", getDestinationSquare("cxb5"), "cxb5")

	Check(t, "h4", getDestinationSquare("Rhxh4+"), "Rhxh4+")
	Check(t, "a5", getDestinationSquare("N3xa5#"), "N3xa5#")
}

func TestGetPieceType(t *testing.T) {
	CheckType(t, PAWN, getPieceType("e4"), "e4")
	CheckType(t, KING, getPieceType("Kd4"), "Kd4")
	CheckType(t, QUEEN, getPieceType("Qdxh1"), "Qdxh1")
	CheckType(t, ROOK, getPieceType("Rxa3"), "Rxa3")
	CheckType(t, KNIGHT, getPieceType("Nf2"), "Nf2")
	CheckType(t, BISHOP, getPieceType("Bxb5"), "Bxb5")
}

func CheckType(t *testing.T, exp int, act int, test string) {
	//fmt.Println("testing, ", test)
	if exp != act {
		t.Error("Testing: " + test + " failed. expected " + strconv.Itoa(exp) + " got:[" + strconv.Itoa(act) + "]")
	}
}

func Check(t *testing.T, exp string, act string, test string) {
	//fmt.Println("testing, ", test)
	if exp != act {
		t.Error("Testing: " + test + " failed. expected " + exp + " got:[" + act + "]")
	}
}
