package battleship

import (
	"bytes"
)

const (
	WATER = iota // Grid has not been shot at yet
	MISS  = iota // Grid was a miss
	HIT   = iota // Grid was a hit
	SHIP  = iota // Grid has an unhit ship in it
)

// A board is a 2D slice of grids, each containing either UNKNOWN|MISS|HIT status codes
type Board [][]int

// Create a 10x10 board and allocate all the splices
func newBoard() (b Board) {
	b = make([][]int, 10)
	for i := range b {
		b[i] = make([]int, 10)
	}

	return
}

// Sets any grid points covered by the given ship to SHIP status
func (b Board) AddShip(s *Ship) {
	if s.Start.X == s.End.X {
		for i := s.Start.Y; i <= s.End.Y; i++ {
			b[s.Start.X][i] = SHIP
		}
	} else {
		for i := s.Start.X; i <= s.End.X; i++ {
			b[i][s.Start.Y] = SHIP
		}
	}
}

/*
String prints the board in a 10x10 grid
~ Water
* Miss
# Ship (unhit)
@ Ship (hit)
*/
func (b Board) String() string {
	var buf bytes.Buffer

	for _, row := range b {
		for _, col := range row {
			switch col {
			case WATER:
				buf.WriteRune('~')
			case MISS:
				buf.WriteRune('*')
			case HIT:
				buf.WriteRune('@')
			case SHIP:
				buf.WriteRune('#')
			}
			buf.WriteRune(' ')
		}
		buf.WriteRune('\n')
	}

	return buf.String()
}
