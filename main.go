package main

import (
	u "dio-solve/util"
	"fmt"
	"time"
)

func main() {
	succCombos := []u.Combo{} // empty slice of elements with Combo type
	numRange := [2]int{-10, 10}
	succCount := 0

	startTime := time.Now().Unix()

	for time.Now().Unix()-startTime < 30 {
		combo := u.Combo{u.RandI(numRange), u.RandI(numRange), u.RandI(numRange)}

		x := float64(combo.X)
		y := float64(combo.Y)
		z := float64(combo.Z)

		// return math.Abs(x-y) == math.Abs(y-z)
		// return math.Pow(x, 2)+math.Pow(y, 2) == math.Pow(z, 2)
		isSucc := x+y == z

		if isSucc && !u.IsDuplicate(succCombos, combo) {
			succCount++
			succCombos = append(succCombos, combo)

			fmt.Println(succCount, combo)
		}

	}

}
