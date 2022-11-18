package score

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
DIE TESTS
*/

// DIE CREATION
func TestDie(t *testing.T) {
	// 1) Sides > Faces
	d1, err := NewDie(6, 4)
	assert.Nil(t, err, "1.1")
	assert.Equal(t, 6, d1.Sides, "1.2")
	assert.Equal(t, 4, d1.Face, "1.3")

	// 2) Faces > Sides
	d2, err := NewDie(4, 6)
	assert.Nil(t, err, "2.1")
	assert.Equal(t, 4, d2.Sides, "2.2")
	assert.Contains(t, []int{1, 2, 3, 4}, d2.Face, "2.3")

	// 3) Faces = 0
	d3, err := NewDie(3, 0)
	assert.Nil(t, err, "3.1")
	assert.Equal(t, 3, d3.Sides, "3.2")
	assert.Contains(t, []int{1, 2, 3}, d3.Face, "3.3")

	// 4) Sides = 0
	d4, err := NewDie(0, 0)
	assert.NotNil(t, err, "4.1")
	_ = d4

	// 5) Sides < 0
	d5, err := NewDie(-1, 3)
	assert.NotNil(t, err, "5.1")
	_ = d5

	// 6) Faces < 0
	d6, err := NewDie(3, -5)
	assert.Nil(t, err, "6.1")
	assert.Equal(t, 3, d6.Sides, "6.2")
	assert.Contains(t, []int{1, 2, 3}, d6.Face, "6.3")
}

// DIE METHODS
func TestDieSetAndGet(t *testing.T) {
	d, err := NewDie(6, 6)
	assert.Nil(t, err, "1.1")
	assert.Equal(t, 6, d.Sides, "1.2")
	assert.Equal(t, 6, d.Face, "1.3")
	assert.Equal(t, d.Face, d.GetFace(), "1.4")

	err = d.SetFace(0)
	assert.NotNil(t, err, "2.1")

	err = d.SetFace(-1)
	assert.NotNil(t, err, "3.1")

	err = d.SetFace(7)
	assert.NotNil(t, err, "4.1")

	for i := 1; i <= d.Sides; i++ {
		d.SetFace(i)
		assert.Equal(t, i, d.Face, "5.1")
		assert.Equal(t, d.Face, d.GetFace(), "5.2")
	}
}

func TestDieRoll(t *testing.T) {
	d, _ := NewDie(6, 6)
	assert.Equal(t, 6, d.Sides, "1.1")
	assert.Equal(t, 6, d.Face, "1.2")

	for i := 0; i < 50; i++ {
		d.Roll()
		assert.Contains(t, []int{1, 2, 3, 4, 5, 6}, d.Face, "2.1")
	}
}

/*
HAND TESTS
*/

// HAND CREATION
func TestHand(t *testing.T) {
	// 1) Qty > 0, Sides > 0
	h1, err := NewHand(3, 6)
	assert.Nil(t, err, "1.1")
	assert.Equal(t, 3, h1.Qty, "1.2")
	assert.Equal(t, 6, h1.Sides, "1.3")

	// 2) Qty = 0, Sides > 0
	h2, err := NewHand(0, 6)
	assert.NotNil(t, err, "2.1")
	_ = h2

	// 3) Qty > 0, Sides = 0
	h3, err := NewHand(3, 0)
	assert.NotNil(t, err, "3.1")
	_ = h3

	// 4) Qty < 0, Sides > 0
	h4, err := NewHand(-2, 6)
	assert.NotNil(t, err, "4.1")
	_ = h4

	// 5) Qty > 0, Sides < 0
	h5, err := NewHand(3, -2)
	assert.NotNil(t, err, "5.1")
	_ = h5
}

// HAND METHODS
func TestHandSetAndGet(t *testing.T) {
	h, err := NewHand(6, 6)
	assert.Nil(t, err, "1.1")
	assert.Equal(t, 6, h.Qty, "1.2")
	assert.Equal(t, 6, h.Sides, "1.3")

	gh := h.GetHand()
	assert.Equal(t, h.Qty, len(gh), "1.4")
	for i := 0; i < h.Qty; i++ {
		assert.Equal(t, h.Dice[i].Face, gh[i], "1.5")
	}

	gh = []int{}
	for i := 0; i < h.Qty; i++ {
		gh = append(gh, rand.Intn(h.Sides)+1)
	}
	assert.Equal(t, h.Qty, len(gh), "1.6")

	err = h.SetHand(gh)
	assert.Nil(t, err, "1.7")
	for i := 0; i < h.Qty; i++ {
		assert.Equal(t, h.Dice[i].Face, gh[i], "1.8")
	}

	h2, err := NewHand(2, 3)
	assert.Nil(t, err, "1.9")
	assert.Equal(t, 6, h.Qty, "1.10")
	assert.Equal(t, 6, h.Sides, "1.11")

	// 2 Dice, 3 Values
	err = h2.SetHand([]int{1, 2, 3})
	assert.NotNil(t, err, "1.12")
	// 2 Dice, 1 Values
	err = h2.SetHand([]int{1})
	assert.NotNil(t, err, "1.13")
	// 3 Sides, 0 Value
	err = h2.SetHand([]int{0, 2})
	assert.NotNil(t, err, "1.14")
	// 3 Sices, 4 Value
	err = h2.SetHand([]int{3, 4})
	assert.NotNil(t, err, "1.15")

}

func TestHandThrow(t *testing.T) {
	// h, _ := NewHand(6, 6)
}
