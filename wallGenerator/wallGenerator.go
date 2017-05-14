package wallGenerator

import (
	"fmt"

	"github.com/Konstantin8105/Convert-INP-to-STD-format/inp"
)

type direction bool

const (
	horizontalDirection direction = false
	verticalDirection             = true
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

	//TODO: add precision
	/*
		// insert points of shell
		n, e, err := createShellByLines(horizLine, vertLine)
		if err != nil {
			return f, err
		}
		e.Name = "Shell"
		//TODO: f.Add(n, e)

		// create vertical shapes
		for i := range vertical {
			n, e, _ := createShapeByLine(vertical[i], vertLine, verticalDirection)
			e.Name = "VertShape"
			//TODO: f.Add(n, e)
		}

		// create horizontal shapes
		for i := range horizontal {
			n, e, _ := createShapeByLine(vertical[i], vertLine, verticalDirection)
			e.Name = "HorizShape"
			//TODO: f.Add(n, e)
		}
	*/
	// debug
	fmt.Println("horizLine = ", horizLine)
	fmt.Println("vertLine = ", vertLine)

	return f, nil
}

func createShapeByLine(shape ShapePosition, line Line, direct direction) (nodes []inp.Node, element inp.Element, err error) {
	/*
		// add point and elements for 01, 12, 13
		sides := [3]sideType{side01, side12, side13}
		for i := range sides {
			s := ([5]side(shape.Sh))[sides[i]]
			if !s.exist {
				continue
			}
			n, e, _ := createShellByLines(s.line, line)

			// coordinates : a, b
		}
	*/
	return nodes, element, nil
}

func createShellByLines(horizLine Line, vertLine Line) (nodes []inp.Node, element inp.Element, err error) {
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
			nodes = append(nodes, n)
			indexOfPoint++
		}
		x := horizSegmengs[j].EndPosition
		n := inp.Node{
			Index: indexOfPoint,
			Coord: [3]float64{x, y, 0.0},
		}
		nodes = append(nodes, n)
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
				nodes = append(nodes, n)
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
			nodes = append(nodes, n)
		}
	}

	// create shells

	shellName := "shell"
	element.Name = shellName
	element.FE, err = inp.GetFiniteElementByName("S4")
	if err != nil {
		return nodes, element, fmt.Errorf("Not correct FE in %v. err = %v", shellName, err)
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
	return nodes, element, nil
}
