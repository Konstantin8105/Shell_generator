package shellGenerator

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkShellWithStiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filename := "testBenchFile.inp"
		// remove file //
		_ = os.Remove(filename)
		// test //
		var shellStiff ShellWithStiffiners
		err := shellStiff.AddShell(Shell{Height: 5.0, Diameter: 1.0, Precition: 0.4})
		if err != nil {
			fmt.Printf("Wrong shell: %v\n", err)
			return

		}
		err = shellStiff.AddStiffiners(Stiffiner{Amount: 5, Height: 0.2, Precition: 0.1})
		if err != nil {
			fmt.Printf("Wrong shell: %v\n", err)
			return
		}

		err = shellStiff.GenerateINP(filename)

		if err != nil {
			fmt.Printf("Wrong mesh: %v\n", err)
			return
		}
		// remove file //
		_ = os.Remove(filename)
	}
}

func Example() {
	filename := "test_shell_cylinder.inp"

	// remove file //
	_ = os.Remove(filename)

	s := Shell{Height: 5, Diameter: 2.0, Precition: 0.2}
	stiffiner := Stiffiner{
		Amount:    6,
		Height:    0.2,
		Precition: 0.5,
	}

	var shellStiff ShellWithStiffiners
	if err := shellStiff.AddShell(s); err != nil {
		fmt.Printf("Wrong shell: %v\n", err)
		return
	}
	if err := shellStiff.AddStiffiners(stiffiner); err != nil {
		fmt.Printf("Wrong stiffiner: %v\n", err)
		return
	}

	err := shellStiff.GenerateINP(filename)
	if err != nil {
		fmt.Printf("Wrong mesh: %v\n", err)
		return
	}

	_ = os.Remove(filename)
}
