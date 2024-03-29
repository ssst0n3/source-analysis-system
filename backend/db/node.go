package db

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/source-analysis-system/model"
)

func listNodesIdByRoot(root uint) (ids []uint, err error) {
	err = DB.Model(model.NodeRelation{}).Select(
		model.SchemaNodeRelation.FieldsByName["Node"].DBName,
	).Where(
		model.SchemaNodeRelation.FieldsByName["Root"].DBName, root,
	).Find(&ids).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func ListNodesByRoot(root uint, withRoot bool) (nodes []model.Node, err error) {
	ids, err := listNodesIdByRoot(root)
	if err != nil {
		return
	}
	err = DB.Model(model.Node{}).Find(&nodes, ids).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if withRoot && root != 0 {
		var rootNode model.Node
		err = DB.Model(model.Node{}).First(&rootNode, root).Error
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		nodes = append(nodes, rootNode)
	}
	return
}

func mapNodesByRoot(root uint) (nodes map[uint]model.Node, err error) {
	nodes = map[uint]model.Node{}
	nodesList, err := ListNodesByRoot(root, false)
	for _, node := range nodesList {
		nodes[node.ID] = node
	}
	return
}

func ListNodeRelationsByRootWithRoot(root uint) (nodeRelations []model.NodeRelation, err error) {
	err = DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Root"].DBName, root).Find(&nodeRelations).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	var rootNodeRelation model.NodeRelation
	err = DB.Model(model.NodeRelation{}).Where(
		model.SchemaNodeRelation.FieldsByName["Node"].DBName, root,
	).First(&rootNodeRelation).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	nodeRelations = append(nodeRelations, rootNodeRelation)
	return
}

func CountNodeRelationsByRootWithRoot(root uint) (count int64, err error) {
	err = DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Root"].DBName, root).Count(&count).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	count += 1
	return
}

func mapNodeRelationsByRoot(root uint) (nodeRelations map[uint]model.NodeRelation, err error) {
	nodeRelations = map[uint]model.NodeRelation{}
	nodeRelationsList, err := ListNodeRelationsByRootWithRoot(root)
	if err != nil {
		return
	}
	for _, nodeRelation := range nodeRelationsList {
		nodeRelations[nodeRelation.Node] = nodeRelation
	}
	return
}

func NodeMatrix(root uint) (matrix [][]model.Node, err error) {
	nodeRelations, err := mapNodeRelationsByRoot(root)
	if err != nil {
		return
	}
	nodes, err := mapNodesByRoot(root)
	if err != nil {
		return
	}
	var rootNode model.Node
	err = DB.Model(model.Node{}).Where("ID", root).First(&rootNode).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	nodes[root] = rootNode
	m := model.NewMatrix(root, nodeRelations)
	m.ChildRecursive(0)
	m.Print()
	matrix = m.DumpNode(nodes)
	return
}

func UpdateMarkdown(id uint, markdown string) (err error) {
	err = DB.Debug().Model(model.Node{}).Where("id", id).Update(
		model.SchemaNode.FieldsByName["Markdown"].DBName, markdown).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func UpdateCollapsed(id uint, collapsed bool) (err error) {
	err = DB.Debug().Model(model.Node{}).Where("id", id).Update(
		model.SchemaNode.FieldsByName["Collapsed"].DBName, collapsed).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
