package app

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestDice_RollOneDice(t *testing.T) {
	dice := newDice()
	var count [6]int
	n := 100000

	for i := 0; i < n; i++ {
		d := dice.RollOneDice()

		assert.GreaterOrEqual(t, d, 1)
		assert.LessOrEqual(t, d, 6)

		switch d {
		case 1:
			count[0]++
		case 2:
			count[1]++
		case 3:
			count[2]++
		case 4:
			count[3]++
		case 5:
			count[4]++
		case 6:
			count[5]++
		default:
			return
		}
	}

	for j := range count {
		assert.GreaterOrEqual(t, count[j], 3*n/6/4)
		assert.LessOrEqual(t, count[j], 5*n/6/4)
	}
}


func TestDice_RollTheDice(t *testing.T) {
	dice := newDice()
	n := 100000
	var count [21]int
	diceCases := [21][]int{
		{1, 1, 1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{2, 2, 2, 2},
		{2, 3},
		{2, 4},
		{2, 5},
		{2, 6},
		{3, 3, 3, 3},
		{3, 4},
		{3, 5},
		{3, 6},
		{4, 4, 4, 4},
		{4, 5},
		{4, 6},
		{5, 5, 5, 5},
		{5, 6},
		{6, 6, 6, 6},
	}

	for i := 0; i < n; i++ {
		d := dice.RollTheDice()

		sort.Ints(d)

		assert.True(t, isDiceCaseValid(diceCases, d))

		for j := range diceCases {
			if isSlicesEqual(d, diceCases[j]) {
				count[j]++
			}
		}
	}

	for k := range count {
		assert.Greater(t, count[k], 0)
	}
}

func isSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func isDiceCaseValid(diceCases [21][]int, d []int) bool {
	for i := range diceCases {
		if isSlicesEqual(diceCases[i], d) {
			return true
		}
	}
	return false
}