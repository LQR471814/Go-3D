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
	From int //? The integer is the index of the vertex
	To   int
}

//Face is the type for all faces no matter how many edges
type Face struct {
	Vertices []int
}

//CreateCube constructs a cube
func CreateCube(x, y, z, sidelen float64) *Mesh {
	vertices := []*Vertex{
		{Position{x - 0.5*sidelen, y - 0.5*sidelen, z - 0.5*sidelen}},

		{Position{x + 0.5*sidelen, y - 0.5*sidelen, z - 0.5*sidelen}},
		{Position{x - 0.5*sidelen, y + 0.5*sidelen, z - 0.5*sidelen}},
		{Position{x - 0.5*sidelen, y - 0.5*sidelen, z + 0.5*sidelen}},

		{Position{x + 0.5*sidelen, y + 0.5*sidelen, z - 0.5*sidelen}},
		{Position{x + 0.5*sidelen, y - 0.5*sidelen, z + 0.5*sidelen}},

		{Position{x - 0.5*sidelen, y + 0.5*sidelen, z + 0.5*sidelen}},

		{Position{x + 0.5*sidelen, y + 0.5*sidelen, z + 0.5*sidelen}},
	}

	edges := []Edge{
		{0, 1},
		{0, 2},
		{0, 3},

		{1, 4},
		{1, 5},

		{2, 4},
		{2, 6},

		{3, 5},
		{3, 6},

		{4, 7},

		{5, 7},

		{6, 7},
	}

	return &Mesh{
		Vertices: vertices,
		Edges:    edges,
	}
}
