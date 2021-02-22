package lib

//Position defines the position type for any coordinate related action
type Position struct {
	X float64
	Y float64
	Z float64
}

//Orientation represents an orientation
type Orientation struct {
	Roll  float64
	Pitch float64
	Yaw   float64
}

//System represents a coordinate system
type System struct {
	Origin   Position
	Rotation Orientation
}

//ToSystem converts a given position to another coordinate system
func ToSystem(s System, p Position) Position {
	return Position{
		(p.X + s.Origin.X),
		(p.Y + s.Origin.Y),
		(p.Z + s.Origin.Z),
	}
}
