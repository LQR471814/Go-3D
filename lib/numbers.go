package lib

import "math"

//Clamp limits the x to a range (use math.NaN to represent no max / min)
func Clamp(x, min, max float64) float64 {
	if max != math.NaN() {
		if x > max {
			return max
		}
	}
	if min != math.NaN() {
		if x < min {
			return min
		}
	}
	return x
}
