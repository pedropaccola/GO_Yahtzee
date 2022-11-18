package game

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"strconv"

	"github.com/pedropaccola/go-yahtzee/score"
)

type Game struct {
	Scoreboard *score.Scoreboard
	Hand       *score.Hand
}

func NewGame() *Game {
	fmt.Println()
	fmt.Println("YAHTZEE") fmt.Println()
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

func (g *Game) Play() {
	for i := 0; i < int(score.LastRule); i++ {
		g.turn()
	}
	fmt.Println("\nCongratulations! You finished the game!")
	fmt.Printf("Total points: %d", g.Scoreboard.TotalPoints())
}

func (g *Game) turn() {
	g.clearScr()
	rolls := 0
	for {
		fmt.Println("Rolling dice...")
		g.Hand.Throw()
		fmt.Println(g.Hand)
		rolls++

		if rolls >= 3 {
			break
		}

	}
}

func (g *Game) reroll() {
	inpInt := []int{}
	for {
		fmt.Println("\nChoose which dice to re-roll")
		fmt.Println("(1, 3, 5) or ('all') or (0) to continue")

		input := ""
		fmt.Scan(&input)

		if strings.ToLower(input) == "all" {
			g.Hand.Throw()
			break
		}

		inpSlice := strings.Fields(input)
		for i, v := range inpSlice {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Sorry, I can't understand")
				continue
			}
			inpInt = append(inpInt, strconv.AtoI)
		}
	}
}

func (g *Game) clearScr() {
	fmt.Println("\033[2J")
	fmt.Println("\033[H")
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func (g *Game) readInput() any {

}
