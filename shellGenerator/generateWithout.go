package shellGenerator

import (
	"fmt"
	"math"

	"github.com/Konstantin8105/Shell_generator/mesh"
)

func (s ShellWithStiffiners) generateWithoutOffset() (m mesh.Mesh, err error) {
	// Plate segment
	var summaryAmountStiffiners int
	for _, st := range s.stiffiners {
		summaryAmountStiffiners += st.Amount
	}
	width := math.Pi * s.shell.Diameter / float64(summaryAmountStiffiners)
	height := math.Min(s.shell.Height, math.Max(s.shell.Precition, s.shell.Height/float64(int(s.shell.Height/s.shell.Precition+0.5))))

	plate, err := separatePlateOnTriangles(mesh.Point{Index: 1, X: 0.0, Y: 0.0, Z: 0.0}, width, height, s.shell.Precition)
	if err != nil {
		return m, fmt.Errorf("Error: %v/n", err)
	}
	// coping by center line axe
	for i := 0; i < summaryAmountStiffiners; i++ {
		angle := 2. * math.Pi / float64(summaryAmountStiffiners) * float64(i)
		for _, point := range plate.Points {
			localAngle := 2.*point.X/s.shell.Diameter + angle
			m.Points = append(m.Points, mesh.Point{
				Index: point.Index + i*len(plate.Points),
				X:     s.shell.Diameter * math.Sin(localAngle),
				Y:     s.shell.Diameter * math.Cos(localAngle),
				Z:     point.Y,
			})
		}
		for _, tr := range plate.Triangles {
			m.Triangles = append(m.Triangles, mesh.Triangle{Indexs: [3]int{
				tr.Indexs[0] + i*len(plate.Points),
				tr.Indexs[1] + i*len(plate.Points),
				tr.Indexs[2] + i*len(plate.Points),
			}})
		}
	}

	return m, nil
}
