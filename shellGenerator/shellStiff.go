package shellGenerator

import (
	"fmt"
	"math"

	"github.com/Konstantin8105/Convert-INP-to-STD-format/inp"
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
	case amountVertStiff < 0:
		return fmt.Errorf("Error: Amount of vertical stiffiners cannot be less zero")
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

	pointsOnLevels := 4
	if s.shell.Precision < s.shell.Diameter {
		pointsOnLevels = int(maxInt(pointsOnLevels, int(math.Pi/math.Asin(s.shell.Precision/s.shell.Diameter)+1)))
	}
	var regularPoints int
	if s.amountVertStiff != 0 {
		regularPoints = int((float64(pointsOnLevels)/float64(s.amountVertStiff) + 0.5)) * s.amountVertStiff
	}
	pointsOnLevels = maxInt(pointsOnLevels, int(regularPoints))

	levels := maxInt(2, int(s.shell.Height/s.shell.Precision+1))
	var regularLevels int
	if s.amountHorizStiff != 0 {
		regularLevels = int((float64(levels)/float64(s.amountHorizStiff+1) + 0.5)) * (s.amountHorizStiff + 1)
	}
	levels = maxInt(levels, regularLevels)

	m, err := s.shell.GenerateMesh(pointsOnLevels, levels)

	// points //
	initPoint := 1 + (levels+1)*pointsOnLevels
	deltaHeight := s.shell.Height / float64(levels)
	deltaLevel := levels / (s.amountHorizStiff + 1)

	for level := 0; level <= levels; level++ {
		elevation := deltaHeight * float64(level)
		if level == levels {
			elevation = s.shell.Height
		}
		for i := 0; i < pointsOnLevels; i++ {
			if (s.amountHorizStiff > 0 && level > 0 && level != levels && float64(level/deltaLevel) == float64(level)/float64(deltaLevel)) ||
				(s.amountVertStiff > 0 && float64(i/(pointsOnLevels/s.amountVertStiff)) == float64(i)/(float64((pointsOnLevels/s.amountVertStiff)))) {
				// add point
				angle := 2. * math.Pi / float64(pointsOnLevels) * float64(i)
				point := inp.Node{
					Index: int(i+pointsOnLevels*level) + initPoint,
					Coord: [3]float64{
						(s.shell.Diameter*0.5 + s.height) * math.Sin(angle),
						elevation,
						(s.shell.Diameter*0.5 + s.height) * math.Cos(angle),
					},
				}
				m.Nodes = append(m.Nodes, point)
			}
		}
	}

	// triangles //
	if s.amountVertStiff > 0 {
		initPoint := 1
		delta := (levels + 1) * pointsOnLevels
		var vertStiff inp.Element
		vertStiff.Name = "VerticalStiffiners"
		vertStiff.FE, err = inp.GetFiniteElementByName("S4")
		if err != nil {
			return err
		}
		for vert := 0; vert < s.amountVertStiff; vert++ {
			for level := 0; level < levels; level++ {
				p1 := initPoint + vert*(pointsOnLevels/s.amountVertStiff) + pointsOnLevels*level
				p2 := initPoint + vert*(pointsOnLevels/s.amountVertStiff) + pointsOnLevels*(level+1)
				p3 := p1 + delta
				p4 := p2 + delta
				quardToRectangle(&vertStiff, p1, p2, p3, p4)
			}
		}
		m.Elements = append(m.Elements, vertStiff)
	}
	if s.amountHorizStiff > 0 {
		initPoint := 1
		delta := (levels + 1) * pointsOnLevels
		var horizStiff inp.Element
		horizStiff.Name = "HorizontalStiffiners"
		horizStiff.FE, err = inp.GetFiniteElementByName("S4")
		if err != nil {
			return err
		}
		for horiz := 0; horiz < s.amountHorizStiff; horiz++ {
			level := (horiz + 1) * deltaLevel
			for i := 0; i < pointsOnLevels; i++ {
				var p1, p2, p3, p4 int
				if i+1 < pointsOnLevels {
					p1 = i + pointsOnLevels*level + 0 + initPoint
					p2 = i + pointsOnLevels*level + 1 + initPoint
					p3 = p1 + delta
					p4 = p2 + delta
				} else {
					p1 = i + pointsOnLevels*level + 0 + initPoint
					p2 = 0 + pointsOnLevels*level + 0 + initPoint
					p3 = p1 + delta
					p4 = p2 + delta
				}
				quardToRectangle(&horizStiff, p1, p2, p3, p4)
			}
		}
		m.Elements = append(m.Elements, horizStiff)
	}

	m.AddUniqueIndexToElements()
	return m.Save(filename)
}
