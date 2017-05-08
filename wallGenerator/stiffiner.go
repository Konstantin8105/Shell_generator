package wallGenerator

// Stiffiner - stiffiner on shell
//     *(L) end of line
//     *
//     *
//     * Line v
//     *
//     *
//     *(0) start position of line
// --------- shell
type Stiffiner Shape

// NewStiffiner - constructor for stiffiner on shell
func NewStiffiner(height, thk float64) (s Stiffiner) {
	var shape Shape
	shape.addSide(side01, height, thk)
	return Stiffiner(shape)
}
