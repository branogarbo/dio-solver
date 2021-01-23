package util

import (
	"math/rand"
)

func RandI(numRange [2]int) int {
	lower := numRange[0]
	upper := numRange[1]

	return rand.Intn(upper-lower+1) + lower
}

func IsDuplicate(slice []Combo, combo Combo) bool {
	for i := 0; i < len(slice); i++ {
		if combo == slice[i] {
			return true
		}
	}

	return false
}
