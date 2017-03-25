package mesh

import (
	"fmt"
	"sort"

	"github.com/Konstantin8105/Convert-INP-to-STD-format/inp"
)

// Point - point of mesh
type Point struct {
	Index   int     // index of point
	X, Y, Z float64 // coordinates
}

type pp []Point

func (a pp) Len() int           { return len(a) }
func (a pp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a pp) Less(i, j int) bool { return a[i].Index < a[j].Index }

// Triangle - elementary triangle with 3 points for mesh
type Triangle struct {
	Indexs [3]int // indexs of 3 points
}

// Mesh - result of mesh generation
type Mesh struct {
	Points    []Point    // points
	Triangles []Triangle // triangles
}

// ConvertInpToMesh - convertor
func (m *Mesh) ConvertInpToMesh(filename string) (err error) {
	// sort points by index
	sort.Sort(pp(m.Points))

	var inpFormat inp.Format
	err = inpFormat.ReadInp(filename)
	if err != nil {
		return err
	}
	fmt.Println("ConvertINP TO MESH = ", m)
	for _, e := range inpFormat.Nodes {
		m.Points = append(m.Points, Point{
			Index: e.Index,
			X:     e.Coord[0],
			Y:     e.Coord[1],
			Z:     e.Coord[2],
		})
	}

	for _, e := range inpFormat.Elements {
		if len(e.IPoint) == 3 && e.ElType == inp.TypeT3D2 {
			m.Triangles = append(m.Triangles, Triangle{Indexs: [3]int{
				e.IPoint[0],
				e.IPoint[1],
				e.IPoint[2],
			}})
		}
	}

	return nil
}
