package main

type Bishop struct {
	Color string
}

func (b *Bishop) PieceType() int {
	return BISHOP
}
func (b *Bishop) ValidMoves(startSquare string) []string {
	return nil
}
func (b *Bishop) GetColor() string {
	return b.Color
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
