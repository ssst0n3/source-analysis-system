package test_db

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/source-analysis-system/model"
	"github.com/ssst0n3/source-analysis-system/test/test_data"
	"gorm.io/gorm"
)

func MakeTableEmpty(db *gorm.DB, modelPtr interface{}) (err error) {
	awesome_reflect.MustPointer(modelPtr)
	err = lightweight_api.DeleteTable(db, modelPtr)
	if err != nil {
		return
	}
	err = lightweight_api.DB.AutoMigrate(modelPtr)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func InitNodes(db *gorm.DB) (err error) {
	return lightweight_api.InitTable(db, &model.Node{}, &test_data.Nodes)
}

func InitNodeRelations(db *gorm.DB) (err error) {
	return lightweight_api.InitTable(db, &model.NodeRelation{}, &test_data.NodeRelations)
}
