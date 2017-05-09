package main

import (
	"github.com/Konstantin8105/Shell_generator/wallGenerator"
)

func main() {
	//cliShellGenerator.Cli()

	h := []wallGenerator.ShapePosition{
		wallGenerator.ShapePosition{wallGenerator.Shape(wallGenerator.NewStiffiner(0.100, 0.010)), 0.3},
	}

	v := []wallGenerator.ShapePosition{
		wallGenerator.ShapePosition{wallGenerator.Shape(wallGenerator.NewAngle(0.120, 0.012, wallGenerator.Type4)), 1.5},
	}

	_, _ = wallGenerator.WallGenerator(wallGenerator.Wall{Lenght: 2.0, Height: 1.0, Thk: 0.005}, v, h)
}
