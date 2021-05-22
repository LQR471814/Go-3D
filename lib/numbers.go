package lib

//Clamp limits the x to a range (use math.NaN to represent no max / min)
func Clamp(x, min, max float64) float64 {
	if x > max {
		return max
	}
	if x < min {
		return min
	}
	return x
}
