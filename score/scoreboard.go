package score 

import "fmt"

type Scoreboard struct {
	Rules []Rule
	Score []int
	Hand *Hand
}

func NewScoreboard() *Scoreboard {
	s := &Scoreboard{
		Rules: make([]Rule, 0, int(LastRule)),
		Score: make([]int, 0, int(LastRule)),
	}
	for i := 0; i < int(LastRule); i++ {
		s.Rules = append(s.Rules, Rule(i)) 
	}
	return s
}

func (s *Scoreboard) Count() int {
	return len(s.Rules)
}

func (s *Scoreboard) GetRule(i int) Rule {
	return s.Rules[i]
}

func (s *Scoreboard) AssignPoints(r Rule) (int, error) {
	row := 0
	for i, v := range s.Rules {
		if v == r {
			row = i
		}
	}
	if s.Score[row] > 0 {
		return 0, fmt.Errorf("Scoreboard for %s already saved", s.Score[row])
	}
	points := r.Points //calculate a rule score
	s.Score[row] = points
	return points, nil
}

func (s *Scoreboard) TotalPoints() int {
	sum := 0
	for _, v := range s.Score {
		sum += v
	}
	return sum
}

func (s *Scoreboard) PointsOverview () {
	strs := []string{}
	for i, v := range s.Rules {
		if 
	}
}