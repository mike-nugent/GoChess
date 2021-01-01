package main

import "reflect"

type Bishop struct {
	Color string
}

func (b *Bishop) PieceType() int {
	return BISHOP
}
func (b *Bishop) ValidMoves(startSquare *Square, board *Board) []Square {

	var line []Square
	line = append(line, board.getLine(b.Color, startSquare, 1, 1)...)
	line = append(line, board.getLine(b.Color, startSquare, -1, 1)...)
	line = append(line, board.getLine(b.Color, startSquare, 1, -1)...)
	line = append(line, board.getLine(b.Color, startSquare, -1, -1)...)

	//TODO - sanity check each square and make sure it's both unique and not nil
	return nil
}
func (b *Bishop) GetColor() string {
	return b.Color
}
func (b *Bishop) ToString() string {
	return b.Color + reflect.TypeOf(b).Name()
}

/*
public List<Square> getSquares(final Square sqr, final Board brd)
{
	List<Square> sqrs = new ArrayList<Square>();
	sqrs.addAll(brd.getLine(this, sqr, 1, 1));
	sqrs.addAll(brd.getLine(this, sqr, -1, 1));
	sqrs.addAll(brd.getLine(this, sqr, 1, -1));
	sqrs.addAll(brd.getLine(this, sqr, -1, -1));

	sqrs = filter(sqrs);

	return brd.getAvailableSquares(this, sqrs);
}
*/
