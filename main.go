package main

import (
	"fmt"

	"github.com/Konstantin8105/Shell_generator/shellGenerator"
)

func main() {
	s := shellGenerator.Shell{Height: 1.5, Diameter: 1.0, Precition: 0.2}
	stiffiner := shellGenerator.Stiffiner{
		Amount:    6,
		Height:    0.2,
		Precition: 0.1,
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

	m, err := shellStiff.Generate(false)
	if err != nil {
		fmt.Printf("Wrong mesh: %v\n", err)
		return
	}
	if err := m.ConvertMeshToINPfile("cylinder.inp"); err != nil {
		fmt.Printf("Wrong mesh: %v\n", err)
		return
	}
}
