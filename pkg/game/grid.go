package game

import (
	"errors"
)

type GridPosition int

const (
	EMPTY  GridPosition = iota // 0
	RED                        // 1
	YELLOW                     // 2
)

type Grid struct {
	_rows int
	_cols int
	_grid [][]GridPosition
}

func NewGrid(rows, cols int) *Grid {
	g := &Grid{_rows: rows, _cols: cols}
	g.InitGrid()
	return g
}

func (g *Grid) InitGrid() {
	g._grid = make([][]GridPosition, g._rows)
	for i := range g._grid {
		g._grid[i] = make([]GridPosition, g._cols)
		for j := range g._grid[i] {
			g._grid[i][j] = EMPTY
		}
	}
}

func (g *Grid) GetGrid() [][]GridPosition {
	return g._grid
}

func (g *Grid) GetColumnCount() int {
	return g._cols
}

func (g *Grid) Print() {
	for _, row := range g._grid {
		for _, cell := range row {
			switch cell {
			case EMPTY:
				print(". ")
			case RED:
				print("R ")
			case YELLOW:
				print("Y ")
			}
		}
		println()
	}
}

func (g *Grid) PlacePiece(col int, piece GridPosition) (int, error) {
	if col < 0 || col >= g._cols {
		return -1, errors.New("invalid column index")
	}

	if piece == EMPTY {
		return -1, errors.New("cannot place an empty piece")
	}

	for i := g._rows - 1; i >= 0; i-- {
		if g._grid[i][col] == EMPTY {
			g._grid[i][col] = piece
			return i, nil
		}
	}

	return -1, errors.New("column is full")
}

func (g *Grid) CheckWin(connectN int, row int, col int, piece GridPosition) bool {
	count := 0

	// Check horizontal
	for c := 0; c < g._cols; c++ {
		if g._grid[row][c] == piece {
			count++
		} else {
			count = 0
		}

		if count == connectN {
			return true
		}
	}

	// Check vertical
	count = 0
	for r := 0; r < g._rows; r++ {
		if g._grid[r][col] == piece {
			count++
		} else {
			count = 0
		}

		if count == connectN {
			return true
		}
	}

	// Check diagonal
	count = 0
	for r := 0; r < g._rows; r++ {
		c := row + col - r
		if c >= 0 && c < g._cols && g._grid[r][c] == piece {
			count++
		} else {
			count = 0
		}

		if count == connectN {
			return true
		}
	}

	// Check anti-diagonal
	count = 0
	for r := 0; r < g._rows; r++ {
		c := col - row + r
		if c >= 0 && c < g._cols && g._grid[r][c] == piece {
			count++
		} else {
			count = 0
		}

		if count == connectN {
			return true
		}
	}

	return false
}
