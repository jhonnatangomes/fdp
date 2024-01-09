package server

type Game struct {
	Players []Player
}

func newGame() *Game {
	return &Game{Players: make([]Player, 0, 4)}
}

func (g *Game) addPlayer(name string) {
	g.Players = append(g.Players, Player{name: name, lives: 3, id: len(g.Players)})
}
