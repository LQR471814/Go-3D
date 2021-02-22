package lib

//Mesh is an object in the 3d world with geometry
type Mesh struct {
	Vertices []*Vertex
	Edges    []Edge
	Faces    []Face
}

//Vertex is the type for all vertices
type Vertex struct {
	Pos Position
}

//Edge is the type for all edges
type Edge struct {
	From *Vertex
	To   *Vertex
}

//Face is the type for all faces no matter how many edges
type Face struct {
	Vertices []*Vertex
}

//CreateCube constructs a cube
func CreateCube(x, y, z, sidelen float64) *Mesh {
	vertices := []*Vertex{
		{Position{x, y, z}},

		{Position{x + sidelen, y, z}},
		{Position{x, y + sidelen, z}},
		{Position{x, y, z + sidelen}},

		{Position{x + sidelen, y + sidelen, z}},
		{Position{x + sidelen, y, z + sidelen}},
		{Position{x + sidelen, y + sidelen, z + sidelen}},

		{Position{x, y + sidelen, z + sidelen}},
	}

	edges := []Edge{
		{vertices[0], vertices[1]},
		{vertices[0], vertices[2]},
		{vertices[0], vertices[3]},

		{vertices[1], vertices[4]},
		{vertices[1], vertices[5]},
		{vertices[1], vertices[6]},

		{vertices[2], vertices[7]},
	}

	return &Mesh{
		Vertices: vertices,
		Edges:    edges,
	}
}
