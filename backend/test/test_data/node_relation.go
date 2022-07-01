package test_data

import (
	"github.com/ssst0n3/source-analysis-system/model"
	"gorm.io/gorm"
)

var (
	NodeRelationTest1_1 = model.NodeRelation{
		Model: gorm.Model{ID: 1},
		Node:  1,
		Root:  0,
		Next:  0,
		Child: 2,
	}
	NodeRelationTest1_2 = model.NodeRelation{
		Model: gorm.Model{ID: 2},
		Node:  2,
		Root:  1,
		Next:  4,
		Child: 3,
	}
	NodeRelationTest1_3 = model.NodeRelation{
		Model: gorm.Model{ID: 3},
		Node:  3,
		Root:  1,
		Next:  0,
		Child: 0,
	}
	NodeRelationTest2_2 = model.NodeRelation{
		Model: gorm.Model{ID: 4},
		Node:  4,
		Root:  1,
		Next:  0,
		Child: 5,
	}
	NodeRelationTest2_3 = model.NodeRelation{
		Model: gorm.Model{ID: 5},
		Node:  5,
		Root:  1,
		Next:  0,
		Child: 0,
	}
	MapNodeRelations = map[uint]model.NodeRelation{
		NodeRelationTest1_2.Node: NodeRelationTest1_2,
		NodeRelationTest1_3.Node: NodeRelationTest1_3,
		NodeRelationTest2_2.Node: NodeRelationTest2_2,
		NodeRelationTest1_1.Node: NodeRelationTest1_1,
		NodeRelationTest2_3.Node: NodeRelationTest2_3,
	}
	NodeRelations = []model.NodeRelation{
		NodeRelationTest1_2,
		NodeRelationTest1_3,
		NodeRelationTest2_2,
		NodeRelationTest1_1,
		NodeRelationTest2_3,
	}
)
