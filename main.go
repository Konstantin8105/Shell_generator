package main

import (
	"fmt"

	"github.com/Konstantin8105/Shell_generator/shellGenerator"
)

func main() {

	//shell()
	//gmshFile()

	s := shellGenerator.Shell{Height: 5, Diameter: 2.0, Precition: 0.2}
	stiffiner := shellGenerator.Stiffiner{
		Amount:    6,
		Height:    0.2,
		Precition: 0.5,
	}

	var shellStiff shellGenerator.ShellWithStiffiners
	if err := shellStiff.AddShell(s); err != nil {
		fmt.Printf("Wrong shell: %v\n", err)
		return
	}
	if err := shellStiff.AddStiffiners(stiffiner); err != nil {
		fmt.Printf("Wrong stiffiner: %v\n", err)
		return
	}

	err := shellStiff.GenerateINP("cylinder.inp")
	if err != nil {
		fmt.Printf("Wrong mesh: %v\n", err)
		return
	}
}

/*
func shell() {
	s := shellGenerator.Shell{Height: 15.0, Diameter: 3.0, Precition: 0.4}
	m, err := s.Generate(true)

	if err != nil {
		fmt.Printf("Wrong mesh: %v\n", err)
		return
	}
	if err := m.ConvertMeshToINPfile("shell.inp"); err != nil {
		fmt.Printf("Wrong mesh: %v\n", err)
		return
	}
}

func gmshFile() {
	filename := "testModel.geo"

	var gh shellGenerator.GmshFormat
	gh.Points = append(gh.Points, shellGenerator.GmshPoint{
		Index:     1,
		X:         0.,
		Y:         0.,
		Z:         0.,
		Precision: 0.1,
	})
	gh.Points = append(gh.Points, shellGenerator.GmshPoint{
		Index:     2,
		X:         1.,
		Y:         0.,
		Z:         0.,
		Precision: 0.1,
	})
	gh.Points = append(gh.Points, shellGenerator.GmshPoint{
		Index:     3,
		X:         0.,
		Y:         1.,
		Z:         0.,
		Precision: 0.5,
	})
	gh.Points = append(gh.Points, shellGenerator.GmshPoint{
		Index:     4,
		X:         1.,
		Y:         -1.,
		Z:         0.,
		Precision: 0.3,
	})

	gh.Lines = append(gh.Lines, shellGenerator.GmshLine{
		Index:           5,
		BeginPointIndex: 2,
		EndPointIndex:   4,
	})

	gh.Arcs = append(gh.Arcs, shellGenerator.GmshArc{
		Index:            6,
		BeginPointIndex:  2,
		CenterPointIndex: 1,
		EndPointIndex:    3,
	})

	gh.Extrudes = append(gh.Extrudes, shellGenerator.GmshExtrude{
		Xextrude:     0,
		Yextrude:     0,
		Zextrude:     3,
		IndexElement: 5,
	})
	gh.Extrudes = append(gh.Extrudes, shellGenerator.GmshExtrude{
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

}*/
