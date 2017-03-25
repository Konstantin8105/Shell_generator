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

		m, err := shellStiff.GenerateINP()

		if err != nil {
			fmt.Printf("Wrong mesh: %v\n", err)
			return
		}
	}
}
