package mesh

import (
	"fmt"
	"sort"

	"github.com/Konstantin8105/Convert-INP-to-STD-format/convertorInp"
)

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

// ConvertMeshToINPfile - convertor
func (m Mesh) ConvertMeshToINPfile(filename string) (err error) {
	err = convertorInp.CreateNewFile(filename, m.convertMeshToINP())
	return err

}

func (m Mesh) convertMeshToINP() (lines []string) {
	lines = append(lines, "*Heading")
	lines = append(lines, " shellGenerator")

	// sort points by index
	sort.Sort(pp(m.Points))

	lines = append(lines, "*NODE")
	for _, point := range m.Points {
		lines = append(lines, fmt.Sprintf("%v, %.10e, %.10e, %.10e", point.Index, point.X, point.Y, point.Z))
	}

	lines = append(lines, "**** ELEMENTS ****")
	lines = append(lines, "*ELEMENT, type=CPS3, ELSET=Shell")
	for i, t := range m.Triangles {
		lines = append(lines, fmt.Sprintf("%v, %v, %v, %v", i+1, t.Indexs[0], t.Indexs[1], t.Indexs[2]))
	}
	return
}
