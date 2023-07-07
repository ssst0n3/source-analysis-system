package db

import (
	"fmt"
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

func TestUpdateChildOrNext(t *testing.T) {
	nodeRelation := model.NodeRelation{
		Node:  99999,
		Root:  99998,
		Next:  10000,
		Child: 10001,
	}
	var count int64

	{
		// prepare
		assert.NoError(t, DB.Model(model.NodeRelation{}).Create(&nodeRelation).Error)
		assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).Count(&count).Error)
		assert.Equal(t, int64(1), count)
	}
	{
		// test
		target := uint(77777)
		assert.NoError(t, UpdateChildOrNext(nodeRelation.Child, target, model.SchemaNodeRelation.FieldsByName["Child"].DBName))
		assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).Count(&count).Error)
		assert.Equal(t, int64(1), count)
		var search model.NodeRelation
		assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).First(&search).Error)
		assert.Equal(t, target, search.Child)
	}

	{
		// clean
		cond := fmt.Sprintf("%s = ?", model.SchemaNodeRelation.FieldsByName["Node"].DBName)
		assert.NoError(t, DB.Where(cond, nodeRelation.Node).Delete(&model.NodeRelation{}).Error)
		assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).Count(&count).Error)
		assert.Equal(t, int64(0), count)
	}
}

func TestHideNode(t *testing.T) {
	nodeRelation := model.NodeRelation{
		Node:  99999,
		Root:  99998,
		Next:  10000,
		Child: 10001,
	}
	var count int64

	{
		// prepare
		assert.NoError(t, DB.Model(model.NodeRelation{}).Create(&nodeRelation).Error)
		assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).Count(&count).Error)
		assert.Equal(t, int64(1), count)
	}

	{
		// test
		assert.NoError(t, HideNode(nodeRelation.Child))
		assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).Count(&count).Error)
		assert.Equal(t, int64(1), count)
		var search model.NodeRelation
		assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).First(&search).Error)
		assert.Equal(t, uint(0), search.Child)
	}

	{
		// clean
		cond := fmt.Sprintf("%s = ?", model.SchemaNodeRelation.FieldsByName["Node"].DBName)
		assert.NoError(t, DB.Where(cond, nodeRelation.Node).Delete(&model.NodeRelation{}).Error)
		assert.NoError(t, DB.Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, nodeRelation.Node).Count(&count).Error)
		assert.Equal(t, int64(0), count)
	}
}
