package players

import (
	"github.com/andrewburian/battleship"
	"log"
	"math/rand"
)

/*
A Random Player simply picks random spots on the board it hasn't
yet shot at, and continues ignorant to any possible feedback.

During placement, it places ships randomly until it find a valid
configuration. (In the least efficient way possible)

Random Player keeps track of its total shots, hits, misses, wins,
and loses. If Logger is not nil, it will log server messages to it.
*/
type RandomPlayer struct {
	Hits, Misses, Shots, Wins, Loses uint
	Logger                           *log.Logger
}

func (p *RandomPlayer) OnHit(g battleship.Grid) {
	p.Hits++
}

func (p *RandomPlayer) OnLoss(g battleship.GameState) {
	p.Loses++
}

func (p *RandomPlayer) OnMiss(g battleship.Grid) {
	p.Loses++
}

func (p *RandomPlayer) OnReceive(g battleship.Grid) {
}

func (p *RandomPlayer) OnSetup(ships []*battleship.Ship) {
	for {
		rnd := rand.Int31()
		for _, ship := range ships {
			switch rnd & 0x1 {
			case 1:
				// up-down
				ship.Start.X = rand.Intn(10)
				ship.Start.Y = rand.Intn(10 - int(ship.Length))
				ship.End.X = ship.Start.X
				ship.End.Y = ship.Start.Y + int(ship.Length) - 1
			case 0:
				// left-right
				ship.Start.Y = rand.Intn(10)
				ship.Start.X = rand.Intn(10 - int(ship.Length))
				ship.End.Y = ship.Start.Y
				ship.End.X = ship.Start.X + int(ship.Length) - 1
			}
			rnd = (rnd >> 1)
		}
		if battleship.ValidateShips(ships) {
			break
		}
	}

}

func (p *RandomPlayer) OnSunk(ship battleship.Ship) {
}

func (p *RandomPlayer) OnTurn(g battleship.GameState) (target battleship.Grid) {
	for {
		target.X = rand.Intn(10)
		target.Y = rand.Intn(10)

		if g.TargetBoard[target.X][target.Y] == battleship.WATER {
			break
		}
	}
	p.Shots++
	return
}

func (p *RandomPlayer) OnWin(g battleship.GameState) {
	p.Wins++
}

func (p *RandomPlayer) OnMessage(msg string) {
	if p.Logger != nil {
		p.Logger.Println(msg)
	}
}
