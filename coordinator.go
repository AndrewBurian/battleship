package battleship

/*
The game coordinator takes 2 games and runs them
*/
type Coordinator struct {
	p1     *Game
	p2     *Game
	obs    Observer
	hasObs bool
}

/*
An observer may be added to a coordinator to get updates after
each shot fired.
This can be for checking rule compliance, rendering, etc
*/
type Observer interface {
	OnChange(GameState, GameState)
}

// Start a new game
func NewGame(p1, p2 Player) *Coordinator {
	return &Coordinator{
		p1: &Game{
			player: p1,
			state: GameState{
				TargetBoard: newBoard(),
				SourceBoard: newBoard(),
			},
		},
		p2: &Game{
			player: p2,
			state: GameState{
				TargetBoard: newBoard(),
				SourceBoard: newBoard(),
			},
		},
	}
}

// Register an observer with the coordinator
func (c *Coordinator) RegisterObserver(obs Observer) {
	c.obs = obs
	c.hasObs = true
}

// Run the game, returns the winner
func (c *Coordinator) Run() Player {

	// Setup the players ships
	c.p1.SetupShips()
	c.p2.SetupShips() //TODO run these in parallel

	// run the game
	var shot Grid
	var resp int
	var sunk *Ship

	for {

		// Player 1 shoots
		shot = c.p1.GetShot()
		resp, sunk = c.p2.ReceiveShot(shot)
		c.p1.ReceiveResponse(shot, resp)
		if sunk != nil {
			c.p1.ReceiveSunk(sunk)
		}

		// notify observer
		if c.hasObs {
			c.obs.OnChange(c.p1.state, c.p2.state)
		}

		// check win/lose
		if c.p2.Lost() {
			c.p1.Win()
			return c.p1.player
		}

		// Player 2 shoots
		shot = c.p2.GetShot()
		resp, sunk = c.p1.ReceiveShot(shot)
		c.p2.ReceiveResponse(shot, resp)
		if sunk != nil {
			c.p2.ReceiveSunk(sunk)
		}

		// notify observer
		if c.hasObs {
			c.obs.OnChange(c.p1.state, c.p2.state)
		}

		// check win/lose
		if c.p1.Lost() {
			c.p2.Win()
			return c.p2.player
		}

	}
}
