package battleship

const (
	UNHIT_SECT = iota // Ship.Hits unhit section
	HIT_SECT   = iota // Ship.Hits hit section
)

// A ship on a game board
type Ship struct {
	Start  Grid   // The starting point of the ship (most top-left)
	End    Grid   // The ending point of the ship (most bottom-right)
	Name   string // Name of the ship
	Length uint   // Total length of the ship
	Hits   []int  // Status of each ship section
	Sunk   bool   // Status of ship sunk
}

// Check if the given ship contains a grid point
func (s *Ship) Contains(g Grid) bool {

	if g.X == s.Start.X {
		if g.Y >= s.Start.Y && g.Y <= s.End.Y {
			return true
		}
	} else if g.Y == s.Start.Y {
		if g.X >= s.Start.X && g.X <= s.End.X {
			return true
		}
	}

	return false
}

// Hit a ship at a given grid point
func (s *Ship) Hit(g Grid) (hit, sunk bool) {

	// check to make sure we contain this grid
	if !s.Contains(g) {

		return false, false
	}

	// mark the appropriate hit
	if g.Y == s.Start.Y {
		s.Hits[g.X-s.Start.X] = HIT_SECT
	} else {
		s.Hits[g.Y-s.Start.Y] = HIT_SECT
	}

	// check if that sunk the ship
	var numhits uint

	for _, h := range s.Hits {
		if h == UNHIT_SECT {
			break
		}
		numhits++
	}
	if numhits == s.Length {
		s.Sunk = true

		return true, true
	}

	// hit, but not sunk

	return true, false
}
