package main

import "fmt"

type Scoreboard struct {
	Rules []Rule
	Score []int
}

func NewScoreboard() 

func (s *Scoreboard) RegisterRules(r []Rule) {
	s.Rules = append(s.Rules, r...)
	s.Score = make([]int, len(s.Rules))
}

func (s *Scoreboard) Count() int {
	return len(s.Rules)
}

func (s *Scoreboard) GetRule(r int) Rule {
	return s.Rules[r]
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