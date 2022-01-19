// Package flyweight 享元模式
package flyweight

var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "車",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
	// ... 其他棋子
}

// ChessPieceUnit 棋子享元
type ChessPieceUnit struct {
	ID    int
	Name  string
	Color string
}

// ChessPiece 棋子
type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

// ChessBoard 棋盘
type ChessBoard struct {
	ChessPieces map[int]*ChessPiece
}

// NewChessPieceUnit 工厂
func NewChessPieceUnit(id int) *ChessPieceUnit {
	return units[id]
}

// NewChessBoard 初始化棋盘
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{
		ChessPieces: map[int]*ChessPiece{},
	}
	for id := range units {
		board.ChessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUnit(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

// Move 移动棋子
func (cb *ChessBoard) Move(id, x, y int) {
	cb.ChessPieces[id].X = x
	cb.ChessPieces[id].Y = y
}
