package shellGenerator

import (
	"fmt"
	"math"

	"github.com/Konstantin8105/Shell_generator/gmsh"
)

// ShellWithStiffiners - input data for shell with stiffiners
type ShellWithStiffiners struct {
	shell            Shell   // data of shell
	amountHorizStiff int     // amount of horizontal stiffiners
	amountVertStiff  int     // amount of vertical stiffiners
	height           float64 // unit - meter. Height of stiffiners
	precision        float64 // unit - meter. Maximal distance between points
}

// AddShell - add shell
func (s *ShellWithStiffiners) AddShell(sh Shell) (err error) {
	if err = sh.check(); err != nil {
		return err
	}
	s.shell = sh
	return nil
}

// AddStiffiners - add stiffiners on shell
func (s *ShellWithStiffiners) AddStiffiners(amountHorizStiff, amountVertStiff int, height, precision float64) (err error) {
	switch {
	case amountHorizStiff < 0:
		return fmt.Errorf("Error: Amount of horizontal stiffiners cannot be less zero")
	case amountVertStiff < 3:
		return fmt.Errorf("Error: Amount of vertical stiffiners cannot be less 3")
	case height <= 0:
		return fmt.Errorf("Error: Height of stiffiners cannot be less or equal zero")
	case precision <= 0:
		return fmt.Errorf("Error: Precision of stiffiners cannot be less or equal zero")
	}
	s.amountHorizStiff = amountHorizStiff
	s.amountVertStiff = amountVertStiff
	s.height = height
	s.precision = precision
	return nil
}

// GenerateINP - Generate mesh of shell
func (s ShellWithStiffiners) GenerateINP(filename string) (err error) {
	err = s.shell.check()
	if err != nil {
		return err
	}

	g, err := s.generateGMSH()
	if err != nil {
		return err
	}

	return g.WriteINP(filename)
}

func (s ShellWithStiffiners) generateGMSH() (g gmsh.Format, err error) {

	// center
	centerPointIndex := 1
	g.AddPoint(gmsh.Point{
		Index:     centerPointIndex,
		X:         0,
		Y:         0,
		Z:         0,
		Precision: s.shell.Precition,
	})

	// cylinder
	if s.amountHorizStiff == 0 {
		startPoint := 10
		startStiffPoint := startPoint + s.amountVertStiff
		startArch := startPoint + s.amountVertStiff*2
		startLine := startPoint + s.amountVertStiff*3
		angleBetweenStiffiners := 2.0 * math.Pi / float64(s.amountVertStiff)

		for i := 0; i < s.amountVertStiff; i++ {
			angle := angleBetweenStiffiners * float64(i)
			// points //
			g.AddPoint(gmsh.Point{
				Index:     startPoint + i,
				X:         s.shell.Diameter / 2. * math.Sin(angle),
				Y:         0.0,
				Z:         s.shell.Diameter / 2. * math.Cos(angle),
				Precision: s.shell.Precition,
			})
			g.AddPoint(gmsh.Point{
				Index:     startStiffPoint + i,
				X:         (s.shell.Diameter/2. + s.height) * math.Sin(angle),
				Y:         0.0,
				Z:         (s.shell.Diameter/2. + s.height) * math.Cos(angle),
				Precision: s.precision,
			})

			// stiffiners //
			g.AddLine(gmsh.Line{
				Index:           startLine + i,
				BeginPointIndex: startPoint + i,
				EndPointIndex:   startStiffPoint + i,
			})

			// arcs //
			if i != s.amountVertStiff-1 {
				g.AddArc(gmsh.Arc{
					Index:            startArch + i,
					BeginPointIndex:  startPoint + i,
					CenterPointIndex: centerPointIndex,
					EndPointIndex:    startPoint + i + 1,
				})
			} else {
				g.AddArc(gmsh.Arc{
					Index:            startArch + i,
					BeginPointIndex:  startPoint + i,
					CenterPointIndex: centerPointIndex,
					EndPointIndex:    startPoint,
				})
			}
		}

		g.ExtrudeAll(0, s.shell.Height, 0)
		return g, nil
	}

	return g, nil
}
