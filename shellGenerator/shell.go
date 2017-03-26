package shellGenerator

import (
	"fmt"
	"math"

	"github.com/Konstantin8105/Shell_generator/mesh"
)

// Shell - input data of shell
type Shell struct {
	Height    float64 // unit - meter. Height of shell
	Diameter  float64 // unit - meter. Diameter of shell
	Precition float64 // unit - meter. Maximal distance between points
}

func (s Shell) check() error {
	switch {
	case s.Height <= 0:
		return fmt.Errorf("Height of shell cannot be less or equal zero")
	case s.Diameter <= 0:
		return fmt.Errorf("Diameter of shell cannot be less or equal zero")
	case s.Precition <= 0:
		return fmt.Errorf("Precition of shell cannot be less or equal zero")
	}
	return nil
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MeshType - type of mesh
type MeshType int

// Types of generated mesh
const (
	RegularMesh MeshType = iota
	OffsetMesh
)

// GenerateMesh - Generate file in Mesh format for shell
func (s Shell) GenerateMesh(mType MeshType, amountOfPointOnLevel, amountLevelsByHeight int) (m mesh.Mesh, err error) {
	err = s.check()
	if err != nil {
		return m, err
	}

	deltaHeigt := s.Height / float64(amountLevelsByHeight-1)

	// init number of point, cannot be less 1
	initPoint := 1

	var iteratorOffset bool
	var angleOffset float64

	for level := 0; level <= amountLevelsByHeight; level++ {
		elevation := deltaHeigt * float64(level)
		if mType == OffsetMesh {
			if iteratorOffset {
				iteratorOffset = false
				angleOffset = 2. * float64(math.Pi) / float64(amountOfPointOnLevel) / 2.
			} else {
				iteratorOffset = true
				angleOffset = 0
			}
		}
		for i := 0; i < amountOfPointOnLevel; i++ {
			angle := 2.*math.Pi/float64(amountOfPointOnLevel)*float64(i) + angleOffset
			m.Points = append(m.Points, mesh.Point{
				Index: int(i+amountOfPointOnLevel*level) + initPoint,
				X:     s.Diameter * math.Sin(angle),
				Y:     elevation,
				Z:     s.Diameter * math.Cos(angle),
			})
		}
	}

	// generate triangles
	iteratorOffset = false
	for level := 0; level < amountLevelsByHeight; level++ {
		if iteratorOffset {
			iteratorOffset = false
		} else {
			iteratorOffset = true
		}
		for i := 0; i < amountOfPointOnLevel; i++ {
			if i+1 < amountOfPointOnLevel {
				m.Triangles = append(m.Triangles, quardToTriangle(
					int(i+amountOfPointOnLevel*level+initPoint),
					int(i+1+amountOfPointOnLevel*level+initPoint),
					int(i+amountOfPointOnLevel*(level+1)+initPoint),
					int(i+1+amountOfPointOnLevel*(level+1)+initPoint),
					iteratorOffset)...)
			} else {
				m.Triangles = append(m.Triangles, quardToTriangle(
					int(i+amountOfPointOnLevel*level+initPoint),
					int(0+amountOfPointOnLevel*level+initPoint),
					int(i+amountOfPointOnLevel*(level+1)+initPoint),
					int(0+amountOfPointOnLevel*(level+1)+initPoint),
					iteratorOffset)...)
			}
		}
	}
	return m, nil
}

// GenerateINP - Generate file in INP format for shell
func (s Shell) GenerateINP(mType MeshType, filename string) (err error) {
	// generate first level of points
	amountOfPointOnLevel := 4
	if s.Precition < s.Diameter {
		amountOfPointOnLevel = int(maxInt(amountOfPointOnLevel, int(math.Pi/math.Asin(s.Precition/s.Diameter)+1)))
	}
	amountLevelsByHeight := maxInt(2, int(s.Height/s.Precition))

	m, err := s.GenerateMesh(mType, amountOfPointOnLevel, amountLevelsByHeight)
	if err != nil {
		return err
	}
	return m.ConvertMeshToINPfile(filename)
}

// Convert 4 points element to 2 triangle
// type = false
//  p3 *---* p4
//     |  /|
//     | / |
//     |/  |
//  p1 *---* p2
//
// type = true
//  p3 *---* p4
//     |\  |
//     | \ |
//     |  \|
//  p1 *---* p2
func quardToTriangle(p1, p2, p3, p4 int, types bool) (t []mesh.Triangle) {
	if types {
		t = append(t, mesh.Triangle{Indexs: [3]int{p1, p3, p2}})
		t = append(t, mesh.Triangle{Indexs: [3]int{p2, p3, p4}})
		return t
	}
	t = append(t, mesh.Triangle{Indexs: [3]int{p1, p4, p2}})
	t = append(t, mesh.Triangle{Indexs: [3]int{p1, p3, p4}})
	return t
}
