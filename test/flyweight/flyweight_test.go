package flyweight

import (
	"testing"

	"github.com/boruns/go-design-pattern/flyweight"
	"github.com/stretchr/testify/assert"
)

func TestFlyWeight(t *testing.T) {

	board1 := flyweight.NewChessBoard()
	board1.Move(1, 1, 2)
	board1.Move(2, 3, 4)
	t.Logf("%v", board1.ChessPieces[1])
	board2 := flyweight.NewChessBoard()
	board2.Move(2, 5, 6)
	board2.Move(1, 7, 8)

	assert.Equal(t, board1.ChessPieces[1].Unit, board2.ChessPieces[1].Unit)
	assert.Equal(t, board1.ChessPieces[2].Unit, board2.ChessPieces[2].Unit)
}
