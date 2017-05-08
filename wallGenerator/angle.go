package wallGenerator

// AngleType - types of angle
type AngleType int

// constants of angle types
const (
	Type1 AngleType = iota
	Type2
	Type3
	Type4
)

// Angle - angle with equal side on shell
//       *          *  *****  *****
//       *          *  *          *
//       *****  *****  *          *
//     -------------------------------- shell
// Type:   1       2      3     4
type Angle Shape

// NewAngle - constructor for angle on shell
func NewAngle(width, thk float64, t AngleType) (a Angle) {
	var shape Shape
	switch t {
	case Type1:
		shape.addSide(side05, width, thk)
		shape.addSide(side01, width, thk)
	case Type2:
		shape.addSide(side04, width, thk)
		shape.addSide(side01, width, thk)
	case Type3:
		shape.addSide(side01, width, thk)
		shape.addSide(side13, width, thk)
	case Type4:
		shape.addSide(side01, width, thk)
		shape.addSide(side12, width, thk)
	}
	return Angle(shape)
}
