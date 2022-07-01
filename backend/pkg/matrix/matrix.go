package matrix

type Matrix struct {
	M         [][]uint
	x         uint
	y         uint
	nextDepth uint
}

func New(root uint) (m Matrix) {
	return Matrix{
		M: [][]uint{
			{root},
		},
	}
}

func (m Matrix) AddChild(node uint) {
	if m.M[m.x+1] == nil {
		m.M = append(m.M, []uint{})
	}
	m.M[m.x+1] = append(m.M[m.x+1], node)
}

func (m Matrix) AddNext(node uint) {
	m.M[m.x] = append(m.M[m.x], node)
}

func (m Matrix) Add(child, next uint) {

}
