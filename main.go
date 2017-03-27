package main

import (
	"fmt"
	"os"

	"github.com/Konstantin8105/Shell_generator/shellGenerator"
)

func main() {
	filename := "test_shell_cylinder.inp"

	// remove file //
	_ = os.Remove(filename)

	s := shellGenerator.Shell{Height: 8, Diameter: 5.0, Precision: 0.3}
	var shellStiff shellGenerator.ShellWithStiffiners
	if err := shellStiff.AddShell(s); err != nil {
		fmt.Printf("Wrong shell: %v\n", err)
		return
	}
	if err := shellStiff.AddStiffiners(3, 9, 0.1, 5); err != nil {
		fmt.Printf("Wrong stiffiner: %v\n", err)
		return
	}

	err := shellStiff.GenerateINP(filename)
	if err != nil {
		fmt.Printf("Wrong mesh: %v\n", err)
		return
	}
}
