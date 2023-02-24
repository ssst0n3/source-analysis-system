package db

import (
	"github.com/ssst0n3/source-analysis-system/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnlinkNodeFromParent(t *testing.T) {
	nodeRelation := model.NodeRelation{
		Node:  99999,
		Root:  99998,
		Next:  10000,
		Child: 10001,
	}
	assert.NoError(t, DB.Model(model.NodeRelation{}).Create(&nodeRelation).Error)
	var count int64
	assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).Count(&count).Error)
	assert.Equal(t, count, int64(1))
	assert.NoError(t, UnlinkNodeFromParent(nodeRelation.Child))
	assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).Count(&count).Error)
	assert.Equal(t, count, int64(0))
}
