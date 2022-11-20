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
	fmt.Println()
	fmt.Println("YAHTZEE")
	fmt.Println()
	fmt.Println("Welcome to the game.")
	fmt.Println("Have fun!.")
	fmt.Println()
	fmt.Println("To exit press [Ctrl + C]")
	fmt.Println()

	for i := 0; i < int(score.LastRule); i++ {
		g.turn()
	}

	fmt.Println("\nCongratulations! You finished the game!")
	fmt.Printf("Total points: %d", g.Scoreboard.TotalPoints())
}

func (g *Game) turn() {
	fmt.Println("\nRolling dice...")
	g.Hand.Throw()

	rolls := 1
	for {
		fmt.Printf("Number of Rolls: %d", rolls)
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
	fmt.Println("\nChoose which dice to re-roll. For example:")
	fmt.Println()
	fmt.Println("(1, 3, 5) separated by commas to re-roll die number 1, 3 and 5.")
	fmt.Println("(all) to re-roll all dice.")
	fmt.Println("(0) to continue without re-roll any dice.")

	for {
		input := ""
		fmt.Scan(&input)

		if len(input) == 0 {
			fmt.Println("Sorry, I can't understand, please repeat input")
			continue
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
			continue
		}
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
