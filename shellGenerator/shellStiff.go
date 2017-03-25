package shellGenerator

import (
	"fmt"
	"math"

	"github.com/Konstantin8105/Shell_generator/gmsh"
)

// Stiffiner - input data of stiffiners
type Stiffiner struct {
	Amount    int     // unit - items. Amount of stiffiners on shell
	Height    float64 // unit - meter. Height of single stiffiner
	Precition float64 // unit - meter. Maximal distance between points
}

func (s Stiffiner) check() (err error) {
	switch {
	case s.Amount < 2:
		return fmt.Errorf("Amount of stiffiners cannot be less 2")
	case s.Height <= 0:
		return fmt.Errorf("Height of stiffiners cannot be less or equal zero")
	case s.Precition <= 0:
		return fmt.Errorf("Precition of stiffiners cannot be less or equal zero")
	}
	return nil
}

// ShellWithStiffiners - input data for shell with stiffiners
type ShellWithStiffiners struct {
	shell      Shell       // data of shell
	stiffiners []Stiffiner // data of stiffiners
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
func (s *ShellWithStiffiners) AddStiffiners(st Stiffiner) (err error) {
	if err = st.check(); err != nil {
		return err
	}
	s.stiffiners = append(s.stiffiners, st)
	return nil
}

// GenerateINP - Generate mesh of shell
func (s ShellWithStiffiners) GenerateINP(filename string) (err error) {
	err = s.shell.check()
	if err != nil {
		return err
	}
	if len(s.stiffiners) < 1 {
		return fmt.Errorf("Error: Please add stiffiners")
	}
	sumStifAmount := 0
	for _, st := range s.stiffiners {
		err = st.check()
		sumStifAmount += st.Amount
		if err != nil {
			return err
		}
	}
	if sumStifAmount < 3 {
		return fmt.Errorf("Error: Minimal amount of stiffiners is 3. You enter %v stiffiners", sumStifAmount)
	}

	gF, err := s.generateGMSH()
	if err != nil {
		return err
	}

	return gF.WriteINP(filename)
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

	sumStifAmount := 0
	for _, st := range s.stiffiners {
		sumStifAmount += st.Amount
	}

	startPoint := 10
	startStiffPoint := startPoint + sumStifAmount
	startArch := startPoint + sumStifAmount*2
	startLine := startPoint + sumStifAmount*3
	angleBetweenStiffiners := 2.0 * math.Pi / float64(sumStifAmount)

	if len(s.stiffiners) > 1 {
		return g, fmt.Errorf("Algorithm error - now implemented only for 1 type of stiffiner. Please connect to developer")
	}

	for i := 0; i < sumStifAmount; i++ {
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
			X:         (s.shell.Diameter/2. + s.stiffiners[0].Height) * math.Sin(angle),
			Y:         0.0,
			Z:         (s.shell.Diameter/2. + s.stiffiners[0].Height) * math.Cos(angle),
			Precision: s.stiffiners[0].Precition,
		})

		// stiffiners //
		g.AddLine(gmsh.Line{
			Index:           startLine + i,
			BeginPointIndex: startPoint + i,
			EndPointIndex:   startStiffPoint + i,
		})

		// arcs //
		if i != sumStifAmount-1 {
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
