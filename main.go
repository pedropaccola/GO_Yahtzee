package main

import (
	"math/rand"
	"time"
  "fmt"

	"github.com/pedropaccola/go-yahtzee/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	g := game.NewGame()
  fmt.Println(g.Scoreboard)
	g.Start()
}
