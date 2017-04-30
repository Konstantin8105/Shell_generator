package imperfection

import (
	"math"

	"github.com/Konstantin8105/Convert-INP-to-STD-format/inp"
)

// Ovalization - type of imperfection. Change all coordinate in model
// amountWave - amount wave, for example, typical ovalization mountWave = 1
// amplitude - amplitude of imperfetion
// angleOffset - offset angle for case with stiffiners may be interest, unit - radiant
func Ovalization(model *inp.Format, amountWave int, amplitude float64, angleOffset float64) {
	if amountWave < 1 {
		amountWave = 1
	}
	for i := range model.Nodes {
		coord := model.Nodes[i].Coord

		// get angle from coordinate
		angle := math.Atan2(coord[0], coord[2])

		amp := amplitude * math.Sin(float64(amountWave*2)*angle+angleOffset)

		dx := amp * math.Sin(angle)
		dz := amp * math.Cos(angle)

		model.Nodes[i].Coord[0] += dx
		model.Nodes[i].Coord[2] += dz
	}
}
