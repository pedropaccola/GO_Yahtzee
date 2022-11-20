package score

import (
	"fmt"
	"strings"
)

type Scoreboard struct {
	Rules  []Rule
	Scores []int
	Hands  []string
}

func NewScoreboard() *Scoreboard {
	s := &Scoreboard{
		Rules:  make([]Rule, 0, int(LastRule)),
		Scores: make([]int, int(LastRule)),
		Hands:  make([]string, int(LastRule)),
	}
	for i := 0; i < int(LastRule); i++ {
		s.Rules = append(s.Rules, Rule(i))
	}
	return s
}

func (s *Scoreboard) String() string {
	strSlice := []string{}
	strSlice = append(strSlice, fmt.Sprintf("=%s=", strings.Repeat("=", 60)))
	strSlice = append(strSlice, fmt.Sprintf("|| %32s %s ||", "SCOREBOARD", strings.Repeat(" ", 23)))
	strSlice = append(strSlice, fmt.Sprintf("=%s=", strings.Repeat("=", 60)))
	strSlice = append(strSlice, fmt.Sprintf("| %-6s | %-20s | %-6s | %-17s |", "Number", "Rule", "Points", "Dice"))
	strSlice = append(strSlice, fmt.Sprintf("+%s+", strings.Repeat("-", 60)))

	for i, rule := range s.Rules {
		str := fmt.Sprintf("| %-6d | %-20s | %-6d | %-17v |", i+1, rule, s.Scores[i], s.Hands[i])
		strSlice = append(strSlice, str)
	}
	strSlice = append(strSlice, fmt.Sprintf("=%s=", strings.Repeat("=", 60)))
	strSlice = append(strSlice, fmt.Sprintf("|| %28s | %-6d | %16s ||", "TOTAL", s.TotalPoints(), " "))
	strSlice = append(strSlice, fmt.Sprintf("=%s=", strings.Repeat("=", 60)))

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

	for i := 0; i < int(LastRule); i++ {
		if i == int(r) {
			s.Scores[i] = rule.Score
			s.Hands[i] = h.GetHandString()
		}
	}

	return rule.Score, nil
}

func (s *Scoreboard) TotalPoints() int {
	sum := 0
	for _, v := range s.Scores {
		sum += v
	}
	return sum
}
