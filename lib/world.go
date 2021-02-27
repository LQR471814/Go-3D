package lib

//Transformable is the interface for all objects that can move, rotate and scale
type Transformable interface {
	Translate()
	Rotate()
	Scale()
}

//Object defines the base object for all objects in the 3d world
type Object struct {
	ObjSys   *System
	Geometry *Mesh
}

//World defines properties for 3d world
type World struct {
	Objects      map[string]*Object
	ActiveCamera *Camera
}

//Camera defines the camera object in the world
type Camera struct {
	Object
	FocalLength float64
}

//Translate moves an object with relative measurements
func (o *Object) Translate(x, y, z float64) {
	o.ObjSys.SetTranslate(
		o.ObjSys.Origin.X+x,
		o.ObjSys.Origin.Y+y,
		o.ObjSys.Origin.Z+z,
	)
}

//Rotate rotates an object with relative measurements
func (o *Object) Rotate(r, p, y float64) {
	o.ObjSys.SetRotation(
		o.ObjSys.Rotation.Roll+r,
		o.ObjSys.Rotation.Pitch+p,
		o.ObjSys.Rotation.Yaw+y,
	)
}

//NewDefaultWorld constructs as default scene and returns a pointer to world
func NewDefaultWorld() *World {
	return &World{
		map[string]*Object{
			"cube": {
				ObjSys: &System{
					Origin:   Position{0, 0, 5},
					Rotation: Orientation{},
				},
				Geometry: CreateCube(0, 0, 0, 2),
			},
		},
		&Camera{
			Object{
				ObjSys: &System{
					Origin:   Position{0, 0, 0},
					Rotation: Orientation{},
				},
				Geometry: &Mesh{},
			},
			2,
		},
	}
}
