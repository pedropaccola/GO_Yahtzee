package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	h := NewHand(5, 6)
	fmt.Println(h.GetHand())
	h.Throw()
	fmt.Println(h.GetHand())
	err := h.Roll([]int{1, 3})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(h.GetHand())

	NewGame()
}
