package battleship

import "fmt"

/*
A game represents a single game board, controlled by a single
player.
*/
type Game struct {
	ships  []*Ship   // List of our ships
	player Player    // Player object making descisions
	state  GameState // State of the boards
}

type GameState struct {
	SourceBoard Board // Status of our board
	TargetBoard Board // Status of enemy board
}

// A grid is a single XY coordinate pair on a game board
type Grid struct {
	X, Y int // grid coordinates using 0 index
}

// Verifies the ships are in a valid arrangement
func ValidateShips(ships []*Ship) bool {
	//TODO
	fmt.Println(ships)
	return true
}

// Gets the status of a grid from a board
func (b Board) Set(g Grid, h int) {

	b[g.X][g.Y] = h

}

// Instructs the player to setup their ships
func (g *Game) SetupShips() {

	// Create all the ships in this game
	var templateShips []*Ship

	templateShips = append(templateShips,
		&Ship{
			Name:   "Aircraft Carrier",
			Length: 5,
			Hits:   make([]int, 5),
		},
		&Ship{
			Name:   "Battleship",
			Length: 4,
			Hits:   make([]int, 4),
		},
		&Ship{
			Name:   "Cruiser",
			Length: 3,
			Hits:   make([]int, 3),
		},
		&Ship{
			Name:   "Submarine",
			Length: 3,
			Hits:   make([]int, 3),
		},
		&Ship{
			Name:   "Destroyer",
			Length: 2,
			Hits:   make([]int, 2),
		},
	)

	for {
		// allow the player to setup the ships
		g.player.OnSetup(templateShips)

		// check valid
		if ValidateShips(templateShips) {
			break
		}

		g.player.OnMessage("Invalid ship arrangement")
	}

	// copy the template ships into the game
	for _, ship := range templateShips {
		g.state.SourceBoard.AddShip(ship)
	}
	g.ships = templateShips

}

// Use the player to produce a shot
func (g *Game) GetShot() Grid {

	// Get player's shot
	return g.player.OnTurn(g.state)
}

// Receive a shot from the other player and send back HIT|MISS
func (g *Game) ReceiveShot(shot Grid) (hitcode int, sunk *Ship) {

	// alert the player
	g.player.OnReceive(shot)

	// check for hits
	for _, ship := range g.ships {
		hit, sunk := ship.Hit(shot)
		if sunk {
			g.state.SourceBoard.Set(shot, HIT)
			return HIT, &Ship{
				Name:   ship.Name,
				Length: ship.Length,
			}
		} else if hit {
			g.state.SourceBoard.Set(shot, HIT)
			return HIT, nil
		}
	}

	g.state.SourceBoard.Set(shot, MISS)
	return MISS, nil
}

// Receive the results from a shot
func (g *Game) ReceiveResponse(shot Grid, resp int) {

	// set the result
	g.state.TargetBoard.Set(shot, resp)

	// send it to the player
	switch resp {
	case HIT:
		g.player.OnHit(shot)
	case MISS:
		g.player.OnMiss(shot)
	}
}

// Receive a report of a sunk ship
func (g *Game) ReceiveSunk(s *Ship) {
	g.player.OnSunk(*s)
}

// Check if the game has been lost
func (g *Game) Lost() bool {

	// check all ships for sunkenness
	for _, ship := range g.ships {
		if !ship.Sunk {
			return false
		}
	}

	// game lost
	g.player.OnLoss(g.state)
	return true
}

// Declare this game the winner
func (g *Game) Win() {
	g.player.OnWin(g.state)
}
