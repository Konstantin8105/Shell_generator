package shellGenerator

import (
	"fmt"
	"math"
)

// Shell - input data of shell
type Shell struct {
	Height    float64 // unit - meter. Height of shell
	Diameter  float64 // unit - meter. Diameter of shell
	Precition float64 // unit - meter. Maximal distance between points
}

func (s Shell) check() error {
	if s.Height <= 0 {
		return fmt.Errorf("Height of shell cannot be less or equal zero")
	}

	if s.Diameter <= 0 {
		return fmt.Errorf("Diameter of shell cannot be less or equal zero")
	}

	if s.Precition <= 0 {
		return fmt.Errorf("Precition of shell cannot be less or equal zero")
	}

	return nil
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Generate mesh of shell
func (s Shell) Generate() (mesh Mesh, err error) {
	err = s.check()
	if err != nil {
		return mesh, err
	}

	// generate first level of points
	amountOfPointOnLevel := 4
	if s.Precition < s.Diameter {
		amountOfPointOnLevel = int(maxInt(amountOfPointOnLevel, int(math.Pi/math.Asin(s.Precition/s.Diameter)+1)))
	}

	amountLevelsByHeight := maxInt(2, int(s.Height/s.Precition))
	deltaHeigt := s.Height / float64(amountLevelsByHeight-1)

	for level := 0; level <= amountLevelsByHeight; level++ {
		for i := 0; i < amountOfPointOnLevel; i++ {
			elevation := deltaHeigt * float64(level)
			angle := 2. * math.Pi / float64(amountOfPointOnLevel) * float64(i)
			mesh.Points = append(mesh.Points, Point{
				index: uint64(i + amountOfPointOnLevel*level),
				X:     s.Diameter * math.Sin(angle),
				Y:     s.Diameter * math.Cos(angle),
				Z:     elevation,
			})
		}
	}

	// generate triangles

	for level := 0; level < amountLevelsByHeight; level++ {
		for i := 0; i < amountOfPointOnLevel; i++ {
			if i+1 < amountOfPointOnLevel {
				mesh.Triangles = append(mesh.Triangles, Triangle{Indexs: [3]uint64{
					uint64(i + amountOfPointOnLevel*level),
					uint64(i + amountOfPointOnLevel*(level+1)),
					uint64(i + 1 + amountOfPointOnLevel*level),
				}})
				mesh.Triangles = append(mesh.Triangles, Triangle{Indexs: [3]uint64{
					uint64(i + 1 + amountOfPointOnLevel*level),
					uint64(i + amountOfPointOnLevel*(level+1)),
					uint64(i + 1 + amountOfPointOnLevel*(level+1)),
				}})
			} else {
				mesh.Triangles = append(mesh.Triangles, Triangle{Indexs: [3]uint64{
					uint64(i + amountOfPointOnLevel*level),
					uint64(i + amountOfPointOnLevel*(level+1)),
					uint64(0 + amountOfPointOnLevel*level),
				}})
				mesh.Triangles = append(mesh.Triangles, Triangle{Indexs: [3]uint64{
					uint64(0 + amountOfPointOnLevel*level),
					uint64(i + amountOfPointOnLevel*(level+1)),
					uint64(0 + amountOfPointOnLevel*(level+1)),
				}})
			}
		}
	}

	return mesh, nil
}
