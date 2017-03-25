package shellGenerator

import (
	"fmt"

	"github.com/Konstantin8105/Shell_generator/mesh"
)

// Plate
//         Width
//    *----------------*
//    |                |
//    |                |
//    |                | Height
//    |                |
//    |                |
// p  *----------------*
//
//    /\ Y
//     |
//     |
//     |
//     *---------->X
//
//
//
// TYPE 1:
// -------
//    *-----*-----*-----* up
//    |    /|    /|    /|
//    |   / |   / |   / |
//    |  /  |  /  |  /  |
//    | /   | /   | /   |
//    |/    |/    |/    |
//    *-----*-----*-----* down
func separatePlateOnTriangles(p mesh.Point, width, height float64, precition float64) (m mesh.Mesh, err error) {
	switch {
	case width <= 0:
		return m, fmt.Errorf("Width of plate is not correct:%v/n", width)
	case height <= 0:
		return m, fmt.Errorf("Height of plate is not correct:%v/n", height)
	case precition <= 0:
		return m, fmt.Errorf("Precition of plate is not correct:%v/n", precition)
	}
	// TYPE 1
	// amount point on down = on up
	amountInternalPoints := int(width/precition + 0.5)
	amountPoints := amountInternalPoints + 2
	distanceBetweenPoints := width / float64(amountPoints-1)
	if distanceBetweenPoints > precition {
		return m, fmt.Errorf("Not correct choosed distance between points: %.3e < %.3e/n", distanceBetweenPoints, precition)
	}

	dx := width / float64(amountPoints-1)
	for i := 0; i <= amountPoints; i++ {
		m.Points = append(m.Points, mesh.Point{
			Index: i + p.Index,
			X:     dx * float64(i),
			Y:     0,
		})
		m.Points = append(m.Points, mesh.Point{
			Index: i + amountPoints + 1 + p.Index,
			X:     dx * float64(i),
			Y:     height,
		})
	}
	for i := 0; i < amountPoints; i++ {
		m.Triangles = append(m.Triangles, quardToTriangle(i+p.Index, i+1+p.Index, i+amountPoints+1+p.Index, i+1+amountPoints+1+p.Index, true)...)
	}
	return m, nil
}

//
// TYPE 2:
// -------
//    *-----*---------* up
//    |    /\        /|
//    |   /  \      / |
//    |  /    \    /  |
//    | /      \  /   |
//    |/        \/    |
//    *----------*----* down
//
//    *-----*---------*----* up
//    |    /\        /\    |
//    |   /  \      /  \   |
//    |  /    \    /    \  |
//    | /      \  /      \ |
//    |/        \/        \|
//    *----------*---------* down
//
// TYPE 3:
// -------
//    *----------*----* up
//    |\        /\    |
//    | \      /  \   |
//    |  \    /    \  |
//    |   \  /      \ |
//    |    \/        \|
//    *-----*---------* down
//
//    *----------*---------* up
//    |\        /\        /|
//    | \      /  \      / |
//    |  \    /    \    /  |
//    |   \  /      \  /   |
//    |    \/        \/    |
//    *-----*---------*----* down
