package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pedropaccola/go-yahtzee/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	g := game.NewGame()

	x := g.Scoreboard.String()
	fmt.Println(x)

	g.Start()
}
