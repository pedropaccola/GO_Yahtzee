package game

import (
	"bufio"
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
	fmt.Println()
	g.Hand.Throw()

	rolls := 1
	for rolls < 3 {
		fmt.Printf("Number of remaining Rolls: %d\n", 3-rolls)
		fmt.Println(g.Hand)

		if !g.reroll() {
			break
		}

		rolls++
	}

	g.score()
}

func (g *Game) reroll() bool {
	inpInt := []int{}
	fmt.Println("Choose which dice to re-roll. For example:")
	fmt.Println()
	fmt.Println("(1, 3, 5) separated by commas to re-roll die number 1, 3 and 5.")
	fmt.Println("(all) to re-roll all dice.")
	fmt.Println("(0) to continue without re-roll any dice.")
	fmt.Println()

InputValidation:
	for {
		inpSli, err := g.readInput()
		if err != nil {
			fmt.Println("Sorry, I can't understand, please repeat input")
			fmt.Println()
			continue
		}
		fmt.Println()

		if len(inpSli) == 0 {
			fmt.Println("Sorry, I can't understand, please repeat input")
			fmt.Println()
			continue
		}

		for _, v := range inpSli {
			if strings.ToLower(v) == "all" {
				fmt.Println("Rolling all dice")
				fmt.Println()
				g.Hand.Throw()
				return true
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Sorry, I can't understand, please repeat input")
				fmt.Println()
				continue InputValidation
			}
			if num == 0 {
				fmt.Println("No re-rolls, continue to scoring")
				fmt.Println()
				return false
			}
			inpInt = append(inpInt, num)
		}

		if err := g.Hand.Roll(inpInt); err != nil {
			fmt.Println(err)
			fmt.Println("Please try again!")
			fmt.Println()
			continue
		}
		fmt.Printf("Rolling dice %v\n", inpInt)
		fmt.Println()
		return true
	}
}

func (g *Game) score() {
	fmt.Println()
	fmt.Println("The Scoreboard:")
	fmt.Println(g.Scoreboard)
	fmt.Println("Your hand:")
	fmt.Println(g.Hand)
	fmt.Println()
	fmt.Printf("Please choose a rule number from 1 to %d to register", score.LastRule)
	fmt.Printf(" your hand scoring.")
	fmt.Println()
	
	for {
		inpSli, err := g.readInput()
		if err != nil {
			fmt.Println("Sorry, I can't understand, please repeat input")
			fmt.Println()
			continue
		}
		if len(inpSli) != 1 {
			fmt.Println("Sorry, I can't understand, please repeat input")
			fmt.Println()
			continue
		}
		inp, err := strconv.Atoi(inpSli[0])
		if err != nil {
			fmt.Println("Sorry, I can't understand, please repeat input")
			fmt.Println()
			continue			
		}
		r := g.Scoreboard.GetRule(inp-1)
		score, err := g.Scoreboard.AssignPoints(r, g.Hand)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Please try again!")
			fmt.Println()
			continue
		}
		fmt.Printf("Registered %d points to the Rule %s", score, r) 
		break
	}
}


// Helper functions
func (g *Game) clearScr() {
	fmt.Println("\033[2J")
	fmt.Println("\033[H")
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func (g *Game) readInput() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	strSli := strings.Fields(strings.Replace(input, ",", " ", -1))
	return strSli, nil
}
