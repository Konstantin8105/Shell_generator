package shellGenerator

import (
	"fmt"
	"sort"
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
func separatePlateOnTriangles(p Point, width, height float64, precition float64) (mesh Mesh, err error) {
	switch {
	case width <= 0:
		return mesh, fmt.Errorf("Width of plate is not correct:%v/n", width)
	case height <= 0:
		return mesh, fmt.Errorf("Height of plate is not correct:%v/n", height)
	case precition <= 0:
		return mesh, fmt.Errorf("Precition of plate is not correct:%v/n", precition)
	}
	// TYPE 1
	// amount point on down = on up
	amountInternalPoints := int(width/precition + 0.5)
	amountPoints := amountInternalPoints + 2
	distanceBetweenPoints := width / float64(amountPoints-1)
	if distanceBetweenPoints > precition {
		return mesh, fmt.Errorf("Not correct choosed distance between points: %.3e < %.3e/n", distanceBetweenPoints, precition)
	}

	dx := width / float64(amountPoints-1)
	for i := 0; i <= amountPoints; i++ {
		mesh.Points = append(mesh.Points, Point{
			index: i + p.index,
			X:     dx * float64(i),
			Y:     0,
		})
		mesh.Points = append(mesh.Points, Point{
			index: i + amountPoints + 1 + p.index,
			X:     dx * float64(i),
			Y:     height,
		})
	}
	// sort points by index
	sort.Sort(pp(mesh.Points))
	for i := 0; i < amountPoints; i++ {
		mesh.Triangles = append(mesh.Triangles, quardToTriangle(i+p.index, i+1+p.index, i+amountPoints+1+p.index, i+1+amountPoints+1+p.index, true)...)
	}
	return mesh, nil
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
