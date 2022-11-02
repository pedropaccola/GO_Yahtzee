package main

import "fmt"

type Game struct {
	Scoreboard *Scoreboard
	Hand       *Hand
}

func NewGame() *Game {
	fmt.Println()
	fmt.Println("YAHTZEE")
	fmt.Println()
	fmt.Println("Welcome to the game. To begin, simply press [Enter]")
	fmt.Println("and follow the instructions on the screen.")
	fmt.Println()
	fmt.Println("To exit press [Ctrl + C]")
	fmt.Println()

	g := &Game{}
	g.Hand = NewHand(5, 6)
	g.Scoreboard = NewScoreboard()
}

func (g *Game) ChooseDiceReroll() {
	for {
		
	}
}