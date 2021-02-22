package lib

import "github.com/veandco/go-sdl2/sdl"

//DrawCircle draws a circle
func DrawCircle(renderer *sdl.Renderer, centreX int, centreY int, radius int) {
	diameter := (radius * 2)

	x := (radius - 1)
	y := 0
	tx := 1
	ty := 1
	error := (tx - diameter)

	for x >= y {
		//  Each of the following renders an octant of the circle
		renderer.DrawPoint(int32(centreX+x), int32(centreY-y))
		renderer.DrawPoint(int32(centreX+x), int32(centreY+y))

		renderer.DrawPoint(int32(centreX-x), int32(centreY-y))
		renderer.DrawPoint(int32(centreX-x), int32(centreY+y))

		renderer.DrawPoint(int32(centreX+y), int32(centreY-x))
		renderer.DrawPoint(int32(centreX+y), int32(centreY+x))

		renderer.DrawPoint(int32(centreX-y), int32(centreY-x))
		renderer.DrawPoint(int32(centreX-y), int32(centreY+x))

		if error <= 0 {
			y++
			error += ty
			ty += 2
		}

		if error > 0 {
			x--
			error += (tx - diameter)
			tx += 2
		}
	}
}

//GetSmaller returns the smaller of two numbers
func GetSmaller(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

//NormalizeScreen normalizes coordinates according to NDC format (ranges 0 - 1)
func NormalizeScreen(w, h, x, y float64) (float64, float64) {
	scale := GetSmaller(float64(w), float64(h))
	return (x * scale / 2) / w, (y * scale / 2) / h
}
