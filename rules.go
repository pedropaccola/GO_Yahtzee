package main

import "sort"

const (
	Aces = iota
	Twos
	Threes
	Fours
	Fives
	Sixes
	ThreeOfAKind
	FourOfAKind
	FullHouse
	SmallStraight
	LargeStraight
	Yahtzee
	Chance
)

type Rule struct {
	Score int
	Name  string
}

func New(h *Hand, rule int) *Rule {
	r := &Rule{}
	switch rule {
	case 0:
		r.Aces(h)

	}
}

func (r *Rule) String() string {
	return r.Name
}

func (r *Rule) Aces(h *Hand) {
	r.Name = "Aces"
	r.UpperSectionScore(1, h)
}

func (r *Rule) Twos(h *Hand) {
	r.Name = "Twos"
	r.UpperSectionScore(2, h)
}

func (r *Rule) Threes(h *Hand) {
	r.Name = "Threes"
	r.UpperSectionScore(3, h)
}

func (r *Rule) Fours(h *Hand) {
	r.Name = "Fours"
	r.UpperSectionScore(4, h)
}

func (r *Rule) Fives() {
	r.Name = "Fives"
	r.UpperSectionScore(5)
}

func (r *Rule) Sixes() {
	r.Name = "Sixes"
	r.UpperSectionScore(6)
}

func (r *Rule) ThreeOfAKind() {
	r.Name = "Three of a Kind"
	r.OfAKind(3)
}

func (r *Rule) FourOfAKind() {
	r.Name = "Four of a Kind"
	r.OfAKind(4)
}

func (r *Rule) FullHouse() {
	r.Name = "Full House"
	m := make(map[int]int)
	for i := 1; i <= 6; i++ {
		m[i]++
	}
	two := false
	three := false
	for _, v := range m {
		if v == 2 {
			two = true
		}
		if v == 3 {
			three = true
		}
	}
	if two == true && three == true {
		r.Score = 25
	}
}

func (r *Rule) SmallStraight() {
	r.Name = "Small Straight"
	l := r.Hands.GetHand()
	sort.Ints(l)
	if r.IsStraight(l[:4]) || r.IsStraight(l[1:]) {
		r.Score = 30
	}
}

func (r *Rule) LargeStraight() {
	r.Name = "Large Straight"
	l := r.Hands.GetHand()
	sort.Ints(l)
	if r.IsStraight(l) {
		r.Score = 40
	}
}

func (r *Rule) Yahtzee() {
	r.Name = "Yahtzee"
	r.OfAKind(5)
}

func (r *Rule) Chance() {
	r.Name = "Chance"
	r.Score = r.Hands.Sum()
}

// Calculation helpers of Upper and Lower Section scorings.
func (r *Rule) UpperSectionScore(v int, h *Hand) {
	r.Score = h.Count(v) * v
}

func (r *Rule) OfAKind(v int) {
	for i := 1; i <= 6; i++ {
		if r.Hands.Count(i) >= v {
			if v != 5 {
				r.Score = r.Hands.Sum()
			} else {
				r.Score = 50
			}
		}
	}
}

func (r *Rule) IsStraight(v []int) bool {
	rep := make(map[int]bool)
	sum := 0
	for _, f := range v {
		if _, value := rep[f]; !value {
			rep[f] = true
			sum += f
		} else {
			return false
		}
	}
	consecutiveSum := len(v) / 2 * (v[0] + v[len(v)-1])

	return sum == consecutiveSum
}
