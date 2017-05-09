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

	// create shell by lines
	// create vertical shapes
	// create horizontal shapes

	// debug
	fmt.Println("horizLine = ", horizLine)
	fmt.Println("vertLine = ", vertLine)

	return f, nil
}
