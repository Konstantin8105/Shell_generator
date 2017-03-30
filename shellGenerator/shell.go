package shellGenerator

import (
	"fmt"
	"math"

	"github.com/Konstantin8105/Convert-INP-to-STD-format/inp"
)

// Shell - input data of shell
type Shell struct {
	Height    float64 // unit - meter. Height of shell
	Diameter  float64 // unit - meter. Diameter of shell
	Precision float64 // unit - meter. Maximal distance between points
}

func (s Shell) check() error {
	switch {
	case s.Height <= 0:
		return fmt.Errorf("Height of shell cannot be less or equal zero")
	case s.Diameter <= 0:
		return fmt.Errorf("Diameter of shell cannot be less or equal zero")
	case s.Precision <= 0:
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
func (s Shell) GenerateMesh(mType MeshType, amountOfPointOnLevel, amountLevelsByHeight int) (m inp.Format, err error) {
	err = s.check()
	if err != nil {
		return m, err
	}
	deltaHeight := s.Height / float64(amountLevelsByHeight)

	// init number of point, cannot be less 1
	initPoint := 1

	var iteratorOffset bool
	var angleOffset float64

	for level := 0; level <= amountLevelsByHeight; level++ {
		elevation := deltaHeight * float64(level)
		if level == amountLevelsByHeight {
			elevation = s.Height
		}
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
			m.Nodes = append(m.Nodes, inp.Node{
				Index: int(i+amountOfPointOnLevel*level) + initPoint,
				Coord: [3]float64{
					s.Diameter * float64(0.5) * math.Sin(angle),
					elevation,
					s.Diameter * float64(0.5) * math.Cos(angle)},
			})
		}
	}

	// generate triangles
	iteratorOffset = false
	var shell inp.Element
	shell.Name = "shell"
	shell.ElType = inp.TypeCPS3
	for level := 0; level < amountLevelsByHeight; level++ {
		if iteratorOffset {
			iteratorOffset = false
		} else {
			iteratorOffset = true
		}
		for i := 0; i < amountOfPointOnLevel; i++ {
			if i+1 < amountOfPointOnLevel {
				quardToTriangle(&shell,
					int(i+amountOfPointOnLevel*level+initPoint),
					int(i+1+amountOfPointOnLevel*level+initPoint),
					int(i+amountOfPointOnLevel*(level+1)+initPoint),
					int(i+1+amountOfPointOnLevel*(level+1)+initPoint),
					iteratorOffset)
			} else {
				quardToTriangle(&shell,
					int(i+amountOfPointOnLevel*level+initPoint),
					int(0+amountOfPointOnLevel*level+initPoint),
					int(i+amountOfPointOnLevel*(level+1)+initPoint),
					int(0+amountOfPointOnLevel*(level+1)+initPoint),
					iteratorOffset)
			}
		}
	}
	m.Elements = append(m.Elements, shell)
	m.AddUniqueIndexToElements()

	m.AddNamedNodesOnLevel(0.0, "Bottom")
	m.AddNamedNodesOnLevel(s.Height, "Top")

	return m, nil
}

// GenerateINP - Generate file in INP format for shell
func (s Shell) GenerateINP(mType MeshType, filename string) (err error) {
	// generate first level of points
	amountOfPointOnLevel := 4
	if s.Precision < s.Diameter {
		amountOfPointOnLevel = int(maxInt(amountOfPointOnLevel, int(math.Pi/math.Asin(s.Precision/s.Diameter)+1)))
	}
	amountLevelsByHeight := maxInt(2, int(s.Height/s.Precision+1))

	m, err := s.GenerateMesh(mType, amountOfPointOnLevel, amountLevelsByHeight)
	if err != nil {
		return err
	}
	return m.Save(filename)
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
func quardToTriangle(element *inp.Element, p1, p2, p3, p4 int, types bool) {
	if types {
		element.Data = append(element.Data, inp.ElementData{
			Index:  -1,
			IPoint: []int{p1, p3, p2},
		})
		element.Data = append(element.Data, inp.ElementData{
			Index:  -1,
			IPoint: []int{p2, p3, p4},
		})
		return
	}
	element.Data = append(element.Data, inp.ElementData{
		Index:  -1,
		IPoint: []int{p1, p4, p2},
	})
	element.Data = append(element.Data, inp.ElementData{
		Index:  -1,
		IPoint: []int{p1, p3, p4},
	})
	return
}
