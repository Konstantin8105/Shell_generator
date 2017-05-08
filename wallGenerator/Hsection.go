package wallGenerator

// Hsection - H or I - section
//    |<--b-->|
//    |       |
//    ********* tf  ---
//        *          |
//        *          |
//        *tw        h
//        *          |
//        *          |
//    ********* tf  ---
type Hsection Shape

// NewHsection - constructor for H or I section
func NewHsection(h, b, tw, tf float64) (a Hsection) {
	var shape Shape
	shape.addSide(side01, h-2.*tf, tw)
	shape.addSide(side12, b/2.0, tf)
	shape.addSide(side13, b/2.0, tf)
	shape.addSide(side04, b/2.0, tf)
	shape.addSide(side05, b/2.0, tf)
	return Hsection(shape)
}
