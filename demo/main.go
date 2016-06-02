package main

import (
	"fmt"
	"github.com/andrewburian/battleship"
	"github.com/andrewburian/battleship/players"
	"math/rand"
	"time"
)

type P1Renderer struct{}

func (o P1Renderer) OnChange(s1, s2 battleship.GameState) {

	fmt.Println(s1.TargetBoard)
	fmt.Println()
	fmt.Println(s1.SourceBoard)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
}

func main() {

	rand.Seed(time.Now().Unix())

	player1 := &players.RandomPlayer{}
	player2 := &players.RandomPlayer{}

	game := battleship.NewGame(player1, player2)

	obs := P1Renderer{}

	game.RegisterObserver(obs)

	game.Run()
}
