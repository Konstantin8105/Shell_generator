package shellGenerator

import (
	"fmt"
	"math"
)

func (s ShellWithStiffiners) generateWithoutOffset() (mesh Mesh, err error) {
	// Plate segment
	var summaryAmountStiffiners int
	for _, st := range s.stiffiners {
		summaryAmountStiffiners += st.Amount
	}
	width := math.Pi * s.shell.Diameter / float64(summaryAmountStiffiners)
	height := math.Min(s.shell.Height, math.Max(s.shell.Precition, s.shell.Height/float64(int(s.shell.Height/s.shell.Precition+0.5))))

	plate, err := separatePlateOnTriangles(Point{1, 0, 0, 0}, width, height, s.shell.Precition)
	if err != nil {
		return mesh, fmt.Errorf("Error: %v/n", err)
	}
	// coping by center line axe
	for i := 0; i < summaryAmountStiffiners; i++ {
		angle := 2. * math.Pi / float64(summaryAmountStiffiners) * float64(i)
		for _, point := range plate.Points {
			localAngle := 2.*point.X/s.shell.Diameter + angle
			newPoint := Point{
				index: point.index + i*len(plate.Points),
				X:     s.shell.Diameter * math.Sin(localAngle),
				Y:     s.shell.Diameter * math.Cos(localAngle),
				Z:     point.Y,
			}
			mesh.Points = append(mesh.Points, newPoint)
		}
		for _, tr := range plate.Triangles {
			mesh.Triangles = append(mesh.Triangles, Triangle{[3]int{
				tr.Indexs[0] + i*len(plate.Points),
				tr.Indexs[1] + i*len(plate.Points),
				tr.Indexs[2] + i*len(plate.Points),
			}})
		}
	}

	return mesh, nil
}
