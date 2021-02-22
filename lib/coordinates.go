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

//Translate sets the position of a pointer to System
func (s *System) Translate(x, y, z float64) {
	s.Origin = Position{x, y, z}
}

//Rotate sets the rotation of a pointer to System
func (s *System) Rotate(r, p, y float64) {
	s.Rotation = Orientation{r, p, y}
}

//ToSystem converts a given position to another coordinate system and is expressed through world coordinates
func ToSystem(s *System, p Position) Position {
	return Position{
		(p.X + s.Origin.X),
		(p.Y + s.Origin.Y),
		(p.Z + s.Origin.Z),
	}
}

//AsRelativeToSystem converts a given world coordinate to what the same world coordinate would be expressed through the new system
func AsRelativeToSystem(s *System, p Position) Position {
	return Position{
		(p.X - s.Origin.X),
		(p.Y - s.Origin.Y),
		(p.Z - s.Origin.Z),
	}
}
