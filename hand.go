package main

import (
	"fmt"
	"math/rand"
)

/*
DIE TYPE AND METHODS
*/

type Die struct {
	Face  int
	Sides int
}

func NewDie(side, face int) *Die {
	d := &Die{
		Sides: side,
		Face:  face,
	}
	if d.Face == 0 {
		d.Roll()
	}
	return d
}

func (d *Die) String() string {
	return fmt.Sprintf("%d", d.Face)
}

func (d *Die) Roll() {
	d.Face = rand.Intn(d.Sides) + 1
}

func (d *Die) SetFace(value int) {
	d.Face = value
}

func (d *Die) GetFace() int {
	return d.Face
}

/*
HAND TYPE AND METHODS
*/

type Hand struct {
	Qty   int
	Sides int
	Dice  []*Die
}

func NewHand(qty, sides int) *Hand {
	h := &Hand{
		Qty:   qty,
		Sides: sides,
		Dice:  make([]*Die, 0, qty),
	}

	for i := 0; i < h.Qty; i++ {
		h.Dice = append(h.Dice, NewDie(sides, 0))
	}

	return h
}

func (h *Hand) String() string {
	s := ""
	for i, v := range h.Dice {
		s += fmt.Sprintf("Die %d has value %d\n", i+1, v.GetFace())
	}
	return s
}

func (h *Hand) Throw() {
	for _, die := range h.Dice {
		die.Roll()
	}
}

func (h *Hand) Roll(dice []int) error {
	for _, v := range dice {
		if v < 1 || v > h.Qty {
			return fmt.Errorf("you only have %d dice", h.Qty)
		}
	}

	for _, v := range dice {
		h.Dice[v-1].Roll()
	}
	return nil
}

func (h *Hand) SetHand(values []int) error {
	if len(values) != h.Qty {
		return fmt.Errorf("you only have %d dice", h.Qty)
	}
	for _, v := range values {
		if v < 1 || v > h.Sides {
			return fmt.Errorf("the dice only have %d sides", h.Sides)
		}
	}
	for i, v := range values {
		h.Dice[i].SetFace(v)
	}

	return nil
}

func (h *Hand) GetHand() []int {
	f := make([]int, 0, h.Qty)
	for _, die := range h.Dice {
		f = append(f, die.GetFace())
	}
	return f
}

func (h *Hand) Count(i int) int {
	s := 0
	for _, v := range h.GetHand() {
		if v == i {
			s++
		}
	}
	return s
}

func (h *Hand) Sum() int {
	s := 0
	for _, v := range h.GetHand() {
		s += v
	}
	return s
}
