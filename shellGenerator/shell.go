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

var ShellName string

func init() {
	ShellName = "shell"
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

// GenerateMesh - Generate file in Mesh format for shell
func (s Shell) GenerateMesh(amountOfPointOnLevel, amountLevelsByHeight int) (m inp.Format, err error) {
	err = s.check()
	if err != nil {
		return m, err
	}
	deltaHeight := s.Height / float64(amountLevelsByHeight)

	// init number of point, cannot be less 1
	initPoint := 1

	var angleOffset float64
	l := (amountLevelsByHeight + 1) * amountOfPointOnLevel
	m.Nodes = make([]inp.Node, 0, l)

	for level := 0; level <= amountLevelsByHeight; level++ {
		elevation := deltaHeight * float64(level)
		if level == amountLevelsByHeight {
			elevation = s.Height
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
	fmt.Println("Generate points")

	// generate triangles
	var shell inp.Element
	shell.Name = ShellName
	shell.FE, err = inp.GetFiniteElementByName("S4")
	if err != nil {
		return m, err
	}
	l2 := amountLevelsByHeight * amountOfPointOnLevel
	shell.Data = make([]inp.ElementData, 0, l2)
	for level := 0; level < amountLevelsByHeight; level++ {
		for i := 0; i < amountOfPointOnLevel; i++ {
			if i+1 < amountOfPointOnLevel {
				quardToTriangle(&shell,
					int(i+amountOfPointOnLevel*level+initPoint),
					int(i+1+amountOfPointOnLevel*level+initPoint),
					int(i+amountOfPointOnLevel*(level+1)+initPoint),
					int(i+1+amountOfPointOnLevel*(level+1)+initPoint))
			} else {
				quardToTriangle(&shell,
					int(i+amountOfPointOnLevel*level+initPoint),
					int(0+amountOfPointOnLevel*level+initPoint),
					int(i+amountOfPointOnLevel*(level+1)+initPoint),
					int(0+amountOfPointOnLevel*(level+1)+initPoint))
			}
		}
	}
	m.Elements = append(m.Elements, shell)
	fmt.Println("Generate elements")
	m.AddUniqueIndexToElements()

	fmt.Println("Return model")
	return m, nil
}

// GenerateINP - Generate file in INP format for shell
func (s Shell) GenerateINP(filename string) (err error) {
	// generate first level of points
	amountOfPointOnLevel := 4
	if s.Precision < s.Diameter {
		amountOfPointOnLevel = int(maxInt(amountOfPointOnLevel, int(math.Pi/math.Asin(s.Precision/s.Diameter)+1)))
	}
	amountLevelsByHeight := maxInt(2, int(s.Height/s.Precision+1))

	m, err := s.GenerateMesh(amountOfPointOnLevel, amountLevelsByHeight)
	if err != nil {
		return err
	}
	return m.Save(filename)
}

// Convert 4 points element to quatric element
//  p3 *---* p4
//     |   |
//     |   |
//     |   |
//  p1 *---* p2
func quardToTriangle(element *inp.Element, p1, p2, p3, p4 int) {
	element.Data = append(element.Data, inp.ElementData{
		Index:  -1,
		IPoint: []int{p1, p2, p4, p3},
	})
	return
}
