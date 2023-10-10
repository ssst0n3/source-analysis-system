package model

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Node struct {
	gorm.Model
	Title           string `json:"title"` // Input from user or calculate from FunctionDeclare
	FunctionDeclare string `json:"function_declare"`
	GitUrl          string `json:"git_url"`
	Note            string `json:"note"`
	Code            string `json:"code"`
	Markdown        string `json:"markdown"`
	Collapsed       bool   `json:"collapsed"`
}

type UpdateMarkdownBody struct {
	Markdown string `json:"markdown"`
}

var SchemaNode schema.Schema

func init() {
	awesome_error.CheckFatal(lightweight_api.InitSchema(&SchemaNode, &Node{}))
}
