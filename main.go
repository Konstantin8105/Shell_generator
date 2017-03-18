package main

import (
	"fmt"

	"github.com/Konstantin8105/Shell_generator/shellGenerator"
)

func main() {
	s := shellGenerator.Shell{Height: 0.5, Diameter: 1.0, Precition: 0.2}
	fmt.Println("Shell = ", s)
	fmt.Println(s.Generate())
}
