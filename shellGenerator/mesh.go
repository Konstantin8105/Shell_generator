package shellGenerator

// Point - point of mesh
type Point struct {
	index   uint64  // index of point
	X, Y, Z float64 // coordinates
}

// Triangle - elementary triangle with 3 points for mesh
type Triangle struct {
	Indexs [3]uint64 // indexs of 3 points
}

// Mesh - result of mesh generation
type Mesh struct {
	Points    []Point    // points
	Triangles []Triangle // triangles
}
