package shellGenerator

import (
	"fmt"
	"os"
)

/*
// Gmsh project created on Mon Mar 20 21:18:09 2017
Point(1) = {0, 0, 0, 1};
Point(2) = {1, 0, 0, 0.1};
Line(1) = {1, 2};
Extrude {0, 2, 0} {
  Line{1};
}
Extrude {4, 0, 0} {
  Line{4};
}
Point(7) = {0, 0, 1, 0.05};
Circle(10) = {2, 1, 7};
Extrude {0, 2, 0} {
  Line{10};
}*/

// GmshPoint - point in Gmsh format
type GmshPoint struct {
	Index     int
	X, Y, Z   float64
	Precision float64
}

func (p GmshPoint) String() string {
	return fmt.Sprintf("Point(%v) = {%.10e, %.10e, %.10e, %.3e};", p.Index, p.X, p.Y, p.Z, p.Precision)
}

// GmshLine - linr in Gmsh format
type GmshLine struct {
	Index           int
	BeginPointIndex int
	EndPointIndex   int
}

func (l GmshLine) String() string {
	return fmt.Sprintf("Line(%v) = {%v , %v};", l.Index, l.BeginPointIndex, l.EndPointIndex)
}

// GmshArc - arc circle in Gmsh format
type GmshArc struct {
	Index            int
	BeginPointIndex  int
	CenterPointIndex int
	EndPointIndex    int
}

func (a GmshArc) String() string {
	return fmt.Sprintf("Circle(%v) = {%v , %v, %v};", a.Index, a.BeginPointIndex, a.CenterPointIndex, a.EndPointIndex)
}

// GmshExtrude - extrude element in Gmsh format
type GmshExtrude struct {
	Xextrude     float64
	Yextrude     float64
	Zextrude     float64
	IndexElement int // index of line or circle arc
}

func (e GmshExtrude) String() string {
	return fmt.Sprintf("Extrude{%v , %v, %v} {\n\tLine{%v};\n}", e.Xextrude, e.Yextrude, e.Zextrude, e.IndexElement)
}

// GmshFormat - complete Gmsh format
type GmshFormat struct {
	Points   []GmshPoint
	Lines    []GmshLine
	Arcs     []GmshArc
	Extrudes []GmshExtrude
}

// Write - write Gmsh file
func (f GmshFormat) Write(filename string) (err error) {
	if len(filename) == 0 {
		return fmt.Errorf("Filename is zero: %v", filename)
	}
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return fmt.Errorf("File %v is exist. Please change the name for saving data", filename)
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Cannot create the file: %v.\nError: %v", filename, err)
	}

	defer func() {
		errFile := file.Close()
		if errFile != nil {
			if err != nil {
				err = fmt.Errorf("%v ; %v", err, errFile)
			} else {
				err = errFile
			}
		}
	}()

	for _, e := range f.Points {
		fmt.Fprintf(file, "%s\n", e)
	}
	for _, e := range f.Lines {
		fmt.Fprintf(file, "%s\n", e)
	}
	for _, e := range f.Arcs {
		fmt.Fprintf(file, "%s\n", e)
	}
	for _, e := range f.Extrudes {
		fmt.Fprintf(file, "%s\n", e)
	}
	return nil
}
