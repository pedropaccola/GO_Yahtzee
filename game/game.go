package game

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/pedropaccola/go-yahtzee/score"
)

type Game struct {
	Scoreboard *score.Scoreboard
	Hand       *score.Hand
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

	h, err := score.NewHand(5, 6)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	s := score.NewScoreboard()

	return &Game{
		Scoreboard: s,
		Hand:       h,
	}
}

func (g *Game) Start() {
	for i := 0; i < int(score.LastRule); i++ {
		g.turn()
	}
	fmt.Println("\nCongratulations! You finished the game!")
	fmt.Printf("Total points: %d", g.Scoreboard.TotalPoints())
}

func (g *Game) turn() {
	g.clearScr()

	fmt.Println("\nRolling dice...")
	g.Hand.Throw()

	rolls := 1
	for {
		fmt.Println(g.Hand)

		if rolls >= 3 {
			break
		}

		if !g.reroll() {
			break
		}

		rolls++
	}

	g.score()
}

func (g *Game) reroll() bool {
	inpInt := []int{}
	fmt.Println("\nChoose which dice to re-roll")
	fmt.Println("(1, 3, 5) or ('all') or (0) to continue")

	input := ""
	fmt.Scan(&input)

	if len(input) == 0 {
		return false
	}

	for _, v := range []byte(input) {
		if string(v) == "0" {
			return false
		}
	}

	if strings.ToLower(input) == "all" {
		g.Hand.Throw()
		return true
	}

	// whole function needs a for loop for error handling
	inpSlice := strings.Fields(input)
	for _, v := range inpSlice {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Sorry, I can't understand")
			fmt.Println("Please try again!")
			continue
		}
		inpInt = append(inpInt, num)
	}
	if err := g.Hand.Roll(inpInt); err != nil {
		fmt.Println(err)
		fmt.Println("Please try again!")
		return false
	}
}

func (g *Game) score() {

}

func (g *Game) clearScr() {
	fmt.Println("\033[2J")
	fmt.Println("\033[H")
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}
