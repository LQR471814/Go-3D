package lib

import (
	"math"
)

//Vector3D represents a collection of 3 numbers
type Vector3D struct {
	X float64
	Y float64
	Z float64
}

//Inverse returns the inverse of Vector3D
func (p Vector3D) Inverse() Vector3D {
	return Vector3D{-p.X, -p.Y, -p.Z}
}

//System represents a coordinate system
type System struct {
	Origin   Vector3D
	Rotation Vector3D
}

//SetTranslate sets the position of a pointer to System
func (s *System) SetTranslate(x, y, z float64) {
	s.Origin = Vector3D{x, y, z}
}

//SetRotation sets the rotation of a pointer to System
func (s *System) SetRotation(r, p, y float64) {
	s.Rotation = Vector3D{r, p, y}
}

//Inverse returns the inverse of the system
func (s *System) Inverse() System {
	return System{s.Origin.Inverse(), s.Rotation.Inverse()}
}

//ToSystem converts a given position to another coordinate system and is expressed through world coordinates,
//this is if the coordinate was the one transforming
func ToSystem(s *System, p Vector3D) Vector3D {
	return ApplySystem(*s, p)
}

//AsRelativeToSystem converts a given world coordinate to what the same world coordinate would be expressed through the new system (This is also the inverse of ToSystem)
//this is if the System is the one transforming
func AsRelativeToSystem(s *System, p Vector3D) Vector3D {
	return ApplySystem(s.Inverse(), p)
}

func ApplySystem(s System, p Vector3D) Vector3D {
	//? On XY plane
	theta := s.Rotation.Z

	X := p.X*math.Cos(theta) - p.Y*math.Sin(theta)
	Y := p.X*math.Sin(theta) + p.Y*math.Cos(theta)

	p.X = X
	p.Y = Y

	//? On XZ plane
	theta = s.Rotation.Y

	X = p.X*math.Cos(theta) - p.Z*math.Sin(theta)
	Z := p.X*math.Sin(theta) + p.Z*math.Cos(theta)

	p.X = X
	p.Z = Z

	//? On YZ plane
	theta = s.Rotation.X

	Y = p.Y*math.Cos(theta) - p.Z*math.Sin(theta)
	Z = p.Y*math.Sin(theta) + p.Z*math.Cos(theta)

	p.Y = Y
	p.Z = Z

	//? Apply Translations
	p.X += s.Origin.X
	p.Y += s.Origin.Y
	p.Z += s.Origin.Z

	return p
}
