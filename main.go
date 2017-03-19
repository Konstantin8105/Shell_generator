package main

import (
	"fmt"

	"github.com/Konstantin8105/Shell_generator/shellGenerator"
)

func main() {
	s := shellGenerator.Shell{Height: 1.5, Diameter: 1.0, Precition: 0.2}
	fmt.Println("Shell = ", s)
	m, err := s.Generate(true)
	if err != nil {
		fmt.Printf("Error in data: %v\n", err)
		return
	}
	m.ConvertMeshToINPfile("cylinder.inp")
}
