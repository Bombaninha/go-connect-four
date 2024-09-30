package game

import (
	"fmt"
)

type Game struct {
	_connectN    int
	Grid         *Grid
	_targetScore int
	Players      []*Player
	_score       map[string]int
}

func NewGame(grid Grid, connectN int, targetScore int) *Game {
	p1 := NewPlayer("Player 1", YELLOW)
	p2 := NewPlayer("Player 2", RED)

	players := []*Player{p1, p2}

	score := make(map[string]int)

	for _, p := range players {
		score[p.GetName()] = 0
	}

	return &Game{
		Grid:         &grid,
		_connectN:    connectN,
		_targetScore: targetScore,
		Players:      players,
		_score:       score,
	}
}

func (g *Game) PrintBoard() {
	g.Grid.Print()
}

func (g *Game) PlayMove(player *Player) (int, int) {
	g.PrintBoard()
	fmt.Println(player.GetName(), "'s turn")
	colCount := g.Grid.GetColumnCount()

	var moveColumn int
	var moveRow int
	var err error

	for {
		fmt.Printf("%s (%d to %d): ", "Enter column number to add a piece", 0, colCount-1)
		fmt.Scan(&moveColumn)

		moveRow, err = g.Grid.PlacePiece(moveColumn, player.GetPieceColor())
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		break
	}

	return moveRow, moveColumn
}

func (g *Game) PlayRound() *Player {
	for {
		for _, player := range g.Players {
			row, col := g.PlayMove(player)
			pieceColor := player.GetPieceColor()

			if g.Grid.CheckWin(g._connectN, row, col, pieceColor) {
				g._score[player.GetName()]++
				return player
			}
		}
	}
}

func (g *Game) Play() {
	maxScore := 0
	var winner *Player

	for maxScore < g._targetScore {
		winner := g.PlayRound()
		fmt.Println("Winner:", winner.GetName())
		maxScore = max(g._score[winner.GetName()], maxScore)

		g.Grid.InitGrid() // reset grid
	}

	print(winner.GetName(), " won the game\n")
}
