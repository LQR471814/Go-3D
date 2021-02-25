package lib

import (
	"math"
)

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

//SetTranslate sets the position of a pointer to System
func (s *System) SetTranslate(x, y, z float64) {
	s.Origin = Position{x, y, z}
}

//SetRotation sets the rotation of a pointer to System
func (s *System) SetRotation(r, p, y float64) {
	s.Rotation = Orientation{r, p, y}
}

//ToSystem converts a given position to another coordinate system and is expressed through world coordinates,
//this is if the coordinate was the one transforming
func ToSystem(s *System, p Position) Position {
	//? On YZ plane
	Y := p.Y*math.Cos(s.Rotation.Roll) - p.Z*math.Sin(s.Rotation.Roll)
	Z := p.Y*math.Sin(s.Rotation.Roll) + p.Z*math.Cos(s.Rotation.Roll)

	p.Y = Y
	p.Z = Z

	//? On XY plane
	X := p.X*math.Cos(s.Rotation.Pitch) - p.Y*math.Sin(s.Rotation.Pitch)
	Y = p.X*math.Sin(s.Rotation.Pitch) + p.Y*math.Cos(s.Rotation.Pitch)

	p.X = X
	p.Y = Y

	//? On XZ plane
	X = p.X*math.Cos(s.Rotation.Yaw) - p.Z*math.Sin(s.Rotation.Yaw)
	Z = p.X*math.Sin(s.Rotation.Yaw) + p.Z*math.Cos(s.Rotation.Yaw)

	p.X = X
	p.Z = Z

	p.X += s.Origin.X
	p.Y += s.Origin.Y
	p.Z += s.Origin.Z

	return p
}

//AsRelativeToSystem converts a given world coordinate to what the same world coordinate would be expressed through the new system (This is also the inverse of ToSystem)
//this is if the System is the one transforming
func AsRelativeToSystem(s *System, p Position) Position {
	//? On YZ plane
	Y := p.Y*math.Cos(-s.Rotation.Roll) - p.Z*math.Sin(-s.Rotation.Roll)
	Z := p.Y*math.Sin(-s.Rotation.Roll) + p.Z*math.Cos(-s.Rotation.Roll)

	p.Y = Y
	p.Z = Z

	//? On XY plane
	X := p.X*math.Cos(-s.Rotation.Pitch) - p.Y*math.Sin(-s.Rotation.Pitch)
	Y = p.X*math.Sin(-s.Rotation.Pitch) + p.Y*math.Cos(-s.Rotation.Pitch)

	p.X = X
	p.Y = Y

	//? On XZ plane
	X = p.X*math.Cos(-s.Rotation.Yaw) - p.Z*math.Sin(-s.Rotation.Yaw)
	Z = p.X*math.Sin(-s.Rotation.Yaw) + p.Z*math.Cos(-s.Rotation.Yaw)

	p.X = X
	p.Z = Z

	p.X -= s.Origin.X
	p.Y -= s.Origin.Y
	p.Z -= s.Origin.Z

	return p
}
