package wallGenerator

type side struct {
	line  Line
	exist bool
}

// SideType - identification of shape
type sideType int

// Sides of universal shape
const (
	side01 sideType = iota
	side12
	side13
	side04
	side05
)

// Shape - universal shape on shell
//
//      2****1****3
//           *
//           *
//           *
//           *
//      4****0****5
// --------------------------- shell (system coordinate: left to right)
type Shape [5]side

// AddSide - add side
func (s *Shape) addSide(st sideType, length, thk float64) {
	(*s)[int(st)].exist = true
	l := make([]LineSegment, 1, 1)
	l[0] = LineSegment{
		BeginPosition: 0.0,
		EndPosition:   length,
		Thk:           thk,
	}
	(*s)[int(st)].line = Line(l)
}
