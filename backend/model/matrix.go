package model

import "fmt"

type Matrix struct {
	matrix        [][]uint
	nodeRelations map[uint]NodeRelation
	x             int
	y             int
}

func NewMatrix(root uint, nodeRelation map[uint]NodeRelation) (m *Matrix) {
	m = &Matrix{
		matrix: [][]uint{
			{root},
		},
		nodeRelations: nodeRelation,
		x:             1,
		y:             1,
	}
	m.NextRecursive(root)
	return
}

func (m *Matrix) NextRecursive(id uint) {
	if id == 0 {
		return
	}
	node := m.nodeRelations[id]
	if node.Next == 0 {
		for y := len(m.matrix[m.x-1]); y < m.y; y++ {
			m.matrix[m.x-1] = append(m.matrix[m.x-1], 0)
		}
		return
	}
	m.matrix[len(m.matrix)-1] = append(m.matrix[len(m.matrix)-1], node.Next)
	newLen := len(m.matrix[len(m.matrix)-1])
	if newLen > m.y {
		for x := 0; x < m.x-1; x++ {
			m.matrix[x] = append(m.matrix[x], 0)
		}
		m.y = newLen
	}
	m.NextRecursive(node.Next)
}

func (m *Matrix) ChildRecursive(x int) {
	for y := m.y - 1; y >= 0; y-- {
		id := m.matrix[x][y]
		if id == 0 {
			continue
		}
		node := m.nodeRelations[id]
		if node.Child == 0 {
			continue
		}
		col := make([]uint, y)
		col = append(col, node.Child)
		m.matrix = append(m.matrix, col)
		m.x += 1
		m.NextRecursive(node.Child)
		m.ChildRecursive(m.x - 1)
	}
}

func (m *Matrix) Print() {
	for y := 0; y <= m.y-1; y++ {
		for x := 0; x <= m.x-1; x++ {
			fmt.Printf("%d ", m.matrix[x][y])
		}
		fmt.Println()
	}
}

func (m *Matrix) DumpNode(nodes map[uint]Node) (matrix [][]Node) {
	for y := 0; y <= m.y-1; y++ {
		var nodeCol []Node
		for x := 0; x <= m.x-1; x++ {
			nodeCol = append(nodeCol, nodes[m.matrix[x][y]])
		}
		matrix = append(matrix, nodeCol)
	}
	return
}
