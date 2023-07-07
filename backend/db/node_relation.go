package db

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/source-analysis-system/model"
)

// UnlinkNodeFromParent DELETE node_relation WHERE child=?
func UnlinkNodeFromParent(child uint) (err error) {
	cond := fmt.Sprintf("%s = ?", model.SchemaNodeRelation.FieldsByName["Child"].DBName)
	err = DB.Debug().Unscoped().Where(cond, child).Delete(&model.NodeRelation{}).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

// UpdateChildOrNext UPDATE node_relation SET child/next = 0 where child/next = ?
func UpdateChildOrNext(old, new uint, field string) (err error) {
	if field != model.SchemaNodeRelation.FieldsByName["Child"].DBName &&
		field != model.SchemaNodeRelation.FieldsByName["Next"].DBName {
		return
	}
	cond := fmt.Sprintf("%s = ?", field)
	err = DB.Debug().Model(&model.NodeRelation{}).Where(cond, old).Update(field, new).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

// HideNode hide node by set child/next filed = 0
func HideNode(id uint) (err error) {
	for _, field := range []string{
		model.SchemaNodeRelation.FieldsByName["Child"].DBName,
		model.SchemaNodeRelation.FieldsByName["Next"].DBName,
	} {
		err = UpdateChildOrNext(id, 0, field)
		if err != nil {
			return
		}
	}
	return
}
