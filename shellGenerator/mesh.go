package shellGenerator

import "github.com/Konstantin8105/Convert-INP-to-STD-format/convertorInp"

// Point - point of mesh
type Point struct {
	index   int     // index of point
	X, Y, Z float64 // coordinates
}

type pp []Point

func (a pp) Len() int           { return len(a) }
func (a pp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a pp) Less(i, j int) bool { return a[i].index < a[j].index }

// Triangle - elementary triangle with 3 points for mesh
type Triangle struct {
	Indexs [3]int // indexs of 3 points
}

// Mesh - result of mesh generation
type Mesh struct {
	Points    []Point    // points
	Triangles []Triangle // triangles
}

// ConvertMeshToINPfile - convertor
func (m Mesh) ConvertMeshToINPfile(filename string) (err error) {
	err = convertorInp.CreateNewFile(filename, m.convertMeshToINP())
	return err
}
