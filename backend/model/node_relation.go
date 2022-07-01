package model

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type NodeRelation struct {
	gorm.Model
	Node  uint `json:"node"`
	Root  uint `json:"root"`
	Next  uint `json:"next"`
	Child uint `json:"child"`
}

var SchemaNodeRelation schema.Schema

func init() {
	awesome_error.CheckFatal(lightweight_api.InitSchema(&SchemaNodeRelation, &NodeRelation{}))
}
