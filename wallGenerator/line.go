package wallGenerator

import (
	"fmt"
)

// LineSegment - segment of line
type LineSegment struct {
	BeginPosition float64
	EndPosition   float64
	Thk           float64
}

// Line - array of position of X
type Line []LineSegment

// NewLine - create new line
func NewLine(lenght float64, thk float64) (line Line) {
	l := make([]LineSegment, 1, 1)
	l[0] = LineSegment{
		BeginPosition: 0.0,
		EndPosition:   lenght,
		Thk:           thk,
	}
	return Line(l)
}

func (l *Line) String() (s string) {
	ls := ([]LineSegment)(*l)
	s += fmt.Sprintf("Line\n")
	for inx := range ls {
		s += fmt.Sprintf("Segment %v\t", inx)
		s += fmt.Sprintf("{%.5e,%5e,%5e}\n", ls[inx].BeginPosition, ls[inx].EndPosition, ls[inx].Thk)
	}
	return s
}

// AddStiffiner - add stiffiner to line
func (l *Line) AddStiffiner(positionX float64, s Stiffiner) (Line, error) {
	_ = s
	return l.AddPoint(positionX)
}

// AddPoint - add point to line
func (l *Line) AddPoint(positionX float64) (Line, error) {
	ls := ([]LineSegment)(*l)
	if positionX < ls[0].BeginPosition {
		return *l, fmt.Errorf("Wrong position less begin")
	}
	if ls[len(ls)-1].EndPosition < positionX {
		return *l, fmt.Errorf("Wrong position more end")
	}
	for inx := range ls {
		segment := ls[inx]
		if float64(segment.BeginPosition) < float64(positionX) && float64(positionX) < float64(segment.EndPosition) {
			s1 := LineSegment{
				BeginPosition: segment.BeginPosition,
				EndPosition:   positionX,
				Thk:           segment.Thk,
			}
			s2 := LineSegment{
				BeginPosition: positionX,
				EndPosition:   segment.EndPosition,
				Thk:           segment.Thk,
			}
			buffer := make([]LineSegment, len(ls)+1, len(ls)+1)
			for i := 0; i < inx; i++ {
				buffer[i] = ls[i]
			}
			buffer[inx] = s1
			buffer[inx+1] = s2
			for i := inx + 1; i < len(ls); i++ {
				buffer[i+1] = ls[i]
			}
			return Line(buffer), nil
		}
	}
	return *l, fmt.Errorf("Point is not on line")
}

// AddShape - add shape to line
func (l *Line) AddShape(positionX float64, shape Shape) (line Line, err error) {
	r := ([5]side(shape))
	line, err = l.AddPoint(positionX)
	if err != nil {
		return line, err
	}
	if shape[side04].exist {
		segments := ([]LineSegment)(r[side04].line)
		for inx := range segments {
			lenght := segments[inx].EndPosition
			line, err = line.AddPoint(positionX - lenght)
			if err != nil {
				return line, err
			}
		}
		line.modifyThickness((shape[side04].line)[0].Thk, positionX-segments[len(segments)-1].EndPosition, positionX)
	}
	if shape[side05].exist {
		segments := ([]LineSegment)(r[side05].line)
		for inx := range segments {
			lenght := segments[inx].EndPosition
			line, err = line.AddPoint(positionX + lenght)
			if err != nil {
				return line, err
			}
		}
		line.modifyThickness((shape[side05].line)[0].Thk, positionX, positionX+segments[len(segments)-1].EndPosition)
	}
	return line, nil
}

func (l *Line) modifyThickness(thk, fromPosition, toPosition float64) {
	segments := ([]LineSegment)(*l)
	for inx := range segments {
		if fromPosition <= segments[inx].BeginPosition && segments[inx].EndPosition <= toPosition {
			segments[inx].Thk += thk
		}
		if toPosition < segments[inx].BeginPosition {
			break
		}
	}
}
