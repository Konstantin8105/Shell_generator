package gmsh_test

import (
	"fmt"
	"os"

	"github.com/Konstantin8105/Shell_generator/gmsh"
)

func Example() {
	filename := "testModel.geo"

	var gh gmsh.Format
	gh.Points = append(gh.Points, gmsh.Point{
		Index:     1,
		X:         0.,
		Y:         0.,
		Z:         0.,
		Precision: 0.1,
	})
	gh.Points = append(gh.Points, gmsh.Point{
		Index:     2,
		X:         1.,
		Y:         0.,
		Z:         0.,
		Precision: 0.1,
	})
	gh.Points = append(gh.Points, gmsh.Point{
		Index:     3,
		X:         0.,
		Y:         1.,
		Z:         0.,
		Precision: 0.5,
	})
	gh.Points = append(gh.Points, gmsh.Point{
		Index:     4,
		X:         1.,
		Y:         -1.,
		Z:         0.,
		Precision: 0.3,
	})

	gh.Lines = append(gh.Lines, gmsh.Line{
		Index:           5,
		BeginPointIndex: 2,
		EndPointIndex:   4,
	})

	gh.Arcs = append(gh.Arcs, gmsh.Arc{
		Index:            6,
		BeginPointIndex:  2,
		CenterPointIndex: 1,
		EndPointIndex:    3,
	})

	gh.Extrudes = append(gh.Extrudes, gmsh.Extrude{
		Xextrude:     0,
		Yextrude:     0,
		Zextrude:     3,
		IndexElement: 5,
	})
	gh.Extrudes = append(gh.Extrudes, gmsh.Extrude{
		Xextrude:     0,
		Yextrude:     0,
		Zextrude:     3,
		IndexElement: 6,
	})

	err := gh.WriteGEO(filename)
	if err != nil {
		fmt.Println("Error in Gmsh format: ", err)
		return
	}

	_ = os.Remove(filename)
}
