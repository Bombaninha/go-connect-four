package game

type Player struct {
	_name       string
	_pieceColor GridPosition
}

func NewPlayer(name string, pieceColor GridPosition) *Player {
	return &Player{_name: name, _pieceColor: pieceColor}
}

func (p *Player) GetName() string {
	return p._name
}

func (p *Player) GetPieceColor() GridPosition {
	return p._pieceColor
}
