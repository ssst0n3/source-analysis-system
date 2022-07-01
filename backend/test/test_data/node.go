package test_data

import (
	"github.com/ssst0n3/source-analysis-system/model"
	"gorm.io/gorm"
)

var (
	NodeTest1_1 = model.Node{
		Model: gorm.Model{
			ID: 1,
		},
		Title: "1-1",
	}
	NodeTest1_2 = model.Node{
		Model: gorm.Model{
			ID: 2,
		},
		Title: "1-2",
	}
	NodeTest1_3 = model.Node{
		Model: gorm.Model{
			ID: 3,
		},
		Title: "1-3",
	}
	NodeTest2_2 = model.Node{
		Model: gorm.Model{
			ID: 4,
		},
		Title: "2-2",
	}
	NodeTest2_3 = model.Node{
		Model: gorm.Model{
			ID: 5,
		},
		Title: "2-3",
	}
	Nodes = []model.Node{
		NodeTest1_1, NodeTest1_2, NodeTest1_3, NodeTest2_2, NodeTest2_3,
	}
)
