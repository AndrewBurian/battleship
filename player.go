package battleship

/*
The Player interface represents the core gameplay logic.

OnSetup is called first, and passed a list of ships with no position information.
Ships should be placed by changing their position, and may also be renamed.

OnTurn is then called every time it is the player's turn to take a shot
*/
type Player interface {
	// On the players turn, player much choose an enemy Grid to fire at
	OnTurn(GameState) Grid

	// After the opposing player's turn, their shot is reported
	OnReceive(Grid)

	// When the player scores a hit, the hit Gird is reported
	OnHit(Grid)

	// When a player misses, the fired on Grid is reported
	OnMiss(Grid)

	// When a player sinks a ship, the sunk ship is reported, but the ship Start and End are not set.
	OnSunk(Ship)

	// On game start, a player must set the Start and End points for all ships
	// This funcion will be called repeatedly if setup is invalid
	OnSetup([]*Ship)

	// If the player wins, this is the final call
	OnWin(GameState)

	// If the player loses, this is the final call
	OnLoss(GameState)

	// Triggered when a message is received from the server
	OnMessage(string)
}
