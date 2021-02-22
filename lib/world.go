package lib

//Object defines the base object for all objects in the 3d world
type Object struct {
	Name     string
	ObjSys   System
	Geometry *Mesh
}

//World defines properties for 3d world
type World struct {
	Objects      []*Object
	ActiveCamera *Camera
}

//Camera defines the camera object in the world
type Camera struct {
	Object
	FocalLength float64
	CamSys      System
}

//NewDefaultWorld constructs as default scene and returns a pointer to world
func NewDefaultWorld() *World {
	return &World{
		[]*Object{
			{
				Name: "cube",
				ObjSys: System{
					Origin:   Position{0, 0, 4},
					Rotation: Orientation{},
				},
				Geometry: CreateCube(0, 0, 0, 2),
			},
		},
		&Camera{
			Object{
				Name: "camera",
				ObjSys: System{
					Origin:   Position{0, 0, 0},
					Rotation: Orientation{},
				},
				Geometry: &Mesh{},
			},
			1,
			System{
				Origin:   Position{0, 0, 0},
				Rotation: Orientation{},
			},
		},
	}
}
