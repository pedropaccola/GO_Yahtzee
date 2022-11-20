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

func (s *Scoreboard) String() string {
	strSlice := []string{}
	strSlice = append(strSlice, fmt.Sprintf("=%s=", strings.Repeat("=", 40)))
	strSlice = append(strSlice, fmt.Sprintf("|| %23s %s ||", "SCOREBOARD", strings.Repeat(" ", 12)))
	strSlice = append(strSlice, fmt.Sprintf("=%s=", strings.Repeat("=", 40)))
	strSlice = append(strSlice, fmt.Sprintf("| %-6s | %-20s | %-6s |", "Number", "Rule", "Points"))
	strSlice = append(strSlice, fmt.Sprintf("+%s+", strings.Repeat("-", 40)))

	for i, v := range s.Rules {
		str := fmt.Sprintf("| %-6d | %-20s | %-6d |", i+1, v, s.Score[i])
		strSlice = append(strSlice, str)
	}
	strSlice = append(strSlice, fmt.Sprintf("=%s=", strings.Repeat("=", 40)))

	output := strings.Join(strSlice, "\n")
	return fmt.Sprintln(output)
}

func (s *Scoreboard) GetRule(i int) Rule {
	return s.Rules[i]
}

func (s *Scoreboard) AssignPoints(r Rule, h *Hand) (int, error) {
	rule, err := NewRule(r, h)
	if err != nil {
		return 0, err
	}

	// row := 0
	// for i, v := range s.Rules {
	// 	if v == r {
	// 		row = i
	// 	}
	// }
	// if s.Score[row] > 0 {
	// 	return 0, fmt.Errorf("scoreboard for %v already saved", s.GetRule(row))
	// }
	// points := 0 // r.Points //calculate a rule score
	// s.Score[row] = points
	return 0, nil
}

func (s *Scoreboard) TotalPoints() int {
	sum := 0
	for _, v := range s.Score {
		sum += v
	}
	return sum
}
