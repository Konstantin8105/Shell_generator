package gmsh

import (
	"fmt"
	"os"
	"os/exec"
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

// Point - point in Gmsh format
type Point struct {
	Index     int
	X, Y, Z   float64
	Precision float64
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%v) = {%.10e, %.10e, %.10e, %.3e};", p.Index, p.X, p.Y, p.Z, p.Precision)
}

// Line - linr in Gmsh format
type Line struct {
	Index           int
	BeginPointIndex int
	EndPointIndex   int
}

func (l Line) String() string {
	return fmt.Sprintf("Line(%v) = {%v , %v};", l.Index, l.BeginPointIndex, l.EndPointIndex)
}

// Arc - arc circle in Gmsh format
type Arc struct {
	Index            int
	BeginPointIndex  int
	CenterPointIndex int
	EndPointIndex    int
}

func (a Arc) String() string {
	return fmt.Sprintf("Circle(%v) = {%v , %v, %v};", a.Index, a.BeginPointIndex, a.CenterPointIndex, a.EndPointIndex)
}

// Extrude - extrude element in Gmsh format
type Extrude struct {
	Xextrude     float64
	Yextrude     float64
	Zextrude     float64
	IndexElement int // index of line or circle arc
}

func (e Extrude) String() string {
	return fmt.Sprintf("Extrude{%v , %v, %v} {\n\tLine{%v};\n}", e.Xextrude, e.Yextrude, e.Zextrude, e.IndexElement)
}

// Format - complete Gmsh format
type Format struct {
	Points   []Point
	Lines    []Line
	Arcs     []Arc
	Extrudes []Extrude
}

// ExtrudeAll - extrude all line and arc
func (f *Format) ExtrudeAll(Xextrude float64, Yextrude float64, Zextrude float64) {
	for _, e := range f.Lines {
		f.AddExtrude(Extrude{
			IndexElement: e.Index,
			Xextrude:     Xextrude,
			Yextrude:     Yextrude,
			Zextrude:     Zextrude,
		})
	}
	for _, e := range f.Arcs {
		f.AddExtrude(Extrude{
			IndexElement: e.Index,
			Xextrude:     Xextrude,
			Yextrude:     Yextrude,
			Zextrude:     Zextrude,
		})
	}
}

// AddPoint - add point
func (f *Format) AddPoint(point Point) {
	f.Points = append(f.Points, point)
}

// AddLine - add line
func (f *Format) AddLine(line Line) {
	f.Lines = append(f.Lines, line)
}

// AddArc - add arc
func (f *Format) AddArc(arch Arc) {
	f.Arcs = append(f.Arcs, arch)
}

// AddExtrude - add extrude
func (f *Format) AddExtrude(e Extrude) {
	f.Extrudes = append(f.Extrudes, e)
}

// WriteGEO - write Gmsh file
func (f Format) WriteGEO(filename string) (err error) {
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

// WriteINP - convertor to mesh format from *.geo GMSH
func (f Format) WriteINP(filename string) (err error) {
	tempFileNameGeo := "temp.geo"

	err = f.WriteGEO(tempFileNameGeo)
	if err != nil {
		return err
	}
	defer func() {
		err2 := os.Remove(tempFileNameGeo)
		if err != nil {
			err = fmt.Errorf("%v\n%v", err, err2)
		} else {
			if err2 != nil {
				err = err2
			}
		}
	}()

	out, err := exec.Command("gmsh", "-0", "-2", "-o", filename, "-format", "inp", tempFileNameGeo).Output()
	if err != nil {
		return fmt.Errorf("ConvertGmsh: %v\n%v", err, out)
	}
	return nil
}
