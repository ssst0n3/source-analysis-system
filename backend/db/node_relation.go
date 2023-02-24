package db

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/source-analysis-system/model"
)

func UnlinkNodeFromParent(child uint) (err error) {
	cond := fmt.Sprintf("%s = ?", model.SchemaNodeRelation.FieldsByName["Child"].DBName)
	err = DB.Debug().Unscoped().Where(cond, child).Delete(&model.NodeRelation{}).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
