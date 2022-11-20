package score

import (
	"fmt"
	"strings"
)

type Scoreboard struct {
	Rules []Rule
	Score []int
}

func NewScoreboard() *Scoreboard {
	s := &Scoreboard{
		Rules: make([]Rule, 0, int(LastRule)),
		Score: make([]int, int(LastRule)),
	}
	for i := 0; i < int(LastRule); i++ {
		s.Rules = append(s.Rules, Rule(i))
	}
	return s
}

// Fix the stringer to print the scoreboard
func (s *Scoreboard) String() string {
	strSlice := []string{}
	strSlice = append(strSlice, "Scoreboard")

	for _, v := range s.Rules {
		strSlice = append(strSlice, v.String())
	}

	output := strings.Join(strSlice, "\n")
	return fmt.Sprintln(output)
}

// func (s *Scoreboard) Count() int {
// 	return len(s.Rules)
// }

// func (s *Scoreboard) GetRule(i int) Rule {
// 	return s.Rules[i]
// }

// func (s *Scoreboard) AssignPoints(r Rule) (int, error) {
// 	row := 0
// 	for i, v := range s.Rules {
// 		if v == r {
// 			row = i
// 		}
// 	}
// 	if s.Score[row] > 0 {
// 		return 0, fmt.Errorf("scoreboard for %v already saved", Rule(s.Score[row]))
// 	}
// 	points := 0 // r.Points //calculate a rule score
// 	s.Score[row] = points
// 	return points, nil
// }

func (s *Scoreboard) TotalPoints() int {
	sum := 0
	for _, v := range s.Score {
		sum += v
	}
	return sum
}

// func (s *Scoreboard) PointsOverview () {
// 	strs := []string{}
// 	for i, v := range s.Rules {
// 		if
// 	}
// }
