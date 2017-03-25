package shellGenerator

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkShellPoint_12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filename := "testBenchFile.inp"
		// remove file //
		_ = os.Remove(filename)
		// test //
		s := Shell{Height: 5.0, Diameter: 1.0, Precition: 2}
		err := s.GenerateINP(true, fileName)

		if err != nil {
			fmt.Printf("Wrong mesh: %v\n", err)
			return
		}
	}
}

func BenchmarkShellPoint_304(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filename := "testBenchFile.inp"
		// remove file //
		_ = os.Remove(filename)
		// test //
		s := Shell{Height: 15.0, Diameter: 1.0, Precition: 0.4}
		err := s.GenerateINP(true, filename)

		if err != nil {
			fmt.Printf("Wrong mesh: %v\n", err)
			return
		}
	}
}

func BenchmarkShellPoint_29704(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filename := "testBenchFile.inp"
		// remove file //
		_ = os.Remove(filename)
		// test //
		s := Shell{Height: 15.0, Diameter: 1.0, Precition: 0.04}
		err := s.GenerateINP(true, filename)

		if err != nil {
			fmt.Printf("Wrong mesh: %v\n", err)
			return
		}
	}
}