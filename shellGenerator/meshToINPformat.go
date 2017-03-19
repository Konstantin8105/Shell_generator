package shellGenerator

import "fmt"

//------------------------------------------
// INP file format
// *Heading
//  cone.inp
// *NODE
// 1, 0, 0, 0
// ******* E L E M E N T S *************
// *ELEMENT, type=T3D2, ELSET=Line1
// 7, 1, 7
// *ELEMENT, type=CPS3, ELSET=Surface17
// 1906, 39, 234, 247
//------------------------------------------

func (m Mesh) convertMeshToINP() (lines []string) {
	lines = append(lines, "*Heading")
	lines = append(lines, " shellGenerator")

	lines = append(lines, "*NODE")
	for _, point := range m.Points {
		lines = append(lines, fmt.Sprintf("%v, %.10e, %.10e, %.10e", point.index, point.X, point.Y, point.Z))
	}

	lines = append(lines, "**** ELEMENTS ****")
	lines = append(lines, "*ELEMENT, type=CPS3, ELSET=Shell")
	for i, t := range m.Triangles {
		lines = append(lines, fmt.Sprintf("%v, %v, %v, %v", i+1, t.Indexs[0], t.Indexs[1], t.Indexs[2]))
	}
	return
}
