package wallGenerator

import (
	"fmt"

	"github.com/Konstantin8105/Convert-INP-to-STD-format/inp"
)

// Wall - basic parameters of wall
type Wall struct {
	Lenght float64
	Height float64
	Thk    float64
}

// ShapePosition - coordinate shapes by position
type ShapePosition struct {
	Sh       Shape
	Position float64
}

// WallGenerator - generate wall with stiffiners
func WallGenerator(wall Wall, vertical []ShapePosition, horizontal []ShapePosition) (f inp.Format, err error) {
	// debug
	fmt.Println("wall = ", wall)
	fmt.Println("horiz = ", horizontal)
	fmt.Println("vert = ", vertical)

	horizLine := NewLine(wall.Lenght, wall.Thk)
	vertLine := NewLine(wall.Height, wall.Thk)

	for i := range vertical {
		horizLine, err = horizLine.AddShape(vertical[i].Position, vertical[i].Sh)
		if err != nil {
			return f, fmt.Errorf("Vertical - %v", err)
		}
	}
	for i := range horizontal {
		vertLine, err = vertLine.AddShape(horizontal[i].Position, horizontal[i].Sh)
		if err != nil {
			return f, fmt.Errorf("Horizontal - %v", err)
		}
	}

	// insert points of shell
	initPointIndex := 1
	indexOfPoint := initPointIndex
	vertSegments := []LineSegment(vertLine)
	horizSegmengs := []LineSegment(horizLine)
	for j := range horizSegmengs {
		y := vertSegments[0].BeginPosition
		if j == 0 {
			x := horizSegmengs[j].BeginPosition
			n := inp.Node{
				Index: indexOfPoint,
				Coord: [3]float64{x, y, 0.0},
			}
			f.Nodes = append(f.Nodes, n)
			indexOfPoint++
		}
		x := horizSegmengs[j].EndPosition
		n := inp.Node{
			Index: indexOfPoint,
			Coord: [3]float64{x, y, 0.0},
		}
		f.Nodes = append(f.Nodes, n)
		indexOfPoint++
	}
	for i := range vertSegments {
		for j := range horizSegmengs {
			y := vertSegments[i].EndPosition
			if j == 0 {
				n := inp.Node{
					Index: indexOfPoint,
					Coord: [3]float64{
						horizSegmengs[j].BeginPosition,
						y,
						0.0,
					},
				}
				indexOfPoint++
				f.Nodes = append(f.Nodes, n)
			}
			x := horizSegmengs[j].EndPosition
			n := inp.Node{
				Index: indexOfPoint,
				Coord: [3]float64{
					x,
					y,
					0.0,
				},
			}
			indexOfPoint++
			f.Nodes = append(f.Nodes, n)
		}
	}

	// create rectangle shell element by points
	shellName := "shell"
	var element inp.Element
	element.Name = shellName
	element.FE, err = inp.GetFiniteElementByName("S4")
	if err != nil {
		return f, fmt.Errorf("Not correct FE in %v. err = %v", shellName, err)
	}
	element.Data = make([]inp.ElementData, len(vertSegments)*len(horizSegmengs), len(vertSegments)*len(horizSegmengs))
	indexOfElement := 1
	for i := range vertSegments {
		for j := range horizSegmengs {
			element.Data[i*len(horizSegmengs)+j] = inp.ElementData{
				Index: indexOfElement,
				IPoint: []int{
					(i+0)*(len(horizSegmengs)+1) + (j + 0) + initPointIndex,
					(i+0)*(len(horizSegmengs)+1) + (j + 1) + initPointIndex,
					(i+1)*(len(horizSegmengs)+1) + (j + 1) + initPointIndex,
					(i+1)*(len(horizSegmengs)+1) + (j + 0) + initPointIndex,
				},
			}
			indexOfElement++
		}
	}
	f.Elements = append(f.Elements, element)
	// create vertical shapes
	// create horizontal shapes

	// debug
	fmt.Println("horizLine = ", horizLine)
	fmt.Println("vertLine = ", vertLine)

	return f, nil
}
