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
		err := shellStiff.AddShell(Shell{Height: 5.0, Diameter: 1.0, Precision: 0.4})
		if err != nil {
			fmt.Printf("Wrong shell: %v\n", err)
			return

		}
		err = shellStiff.AddStiffiners(0, 5, 0.2, 0.1)
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

	s := Shell{Height: 5, Diameter: 2.0, Precision: 0.2}
	var shellStiff ShellWithStiffiners
	if err := shellStiff.AddShell(s); err != nil {
		fmt.Printf("Wrong shell: %v\n", err)
		return
	}
	if err := shellStiff.AddStiffiners(0, 6, 0.2, 0.5); err != nil {
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
