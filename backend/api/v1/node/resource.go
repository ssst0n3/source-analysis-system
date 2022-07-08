package node

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/source-analysis-system/db"
	"github.com/ssst0n3/source-analysis-system/model"
	"net/http"
	"strconv"
)

const ResourceName = "node"

var Resource = lightweight_api.NewResource(ResourceName, model.SchemaNode.Table, model.Node{}, "")

func ParseId(c *gin.Context) (id uint64, err error) {
	rootId := c.Param("id")
	id, err = strconv.ParseUint(rootId, 10, 64)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	return
}

func Matrix(c *gin.Context) {
	id, err := ParseId(c)
	if err != nil {
		return
	}
	matrix, err := db.NodeMatrix(uint(id))
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	c.JSON(http.StatusOK, matrix)
}

func UpdateMarkdown(c *gin.Context) {
	id, err := ParseId(c)
	if err != nil {
		return
	}
	var body model.UpdateMarkdownBody
	err = c.BindJSON(&body)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	err = db.UpdateMarkdown(uint(id), body.Markdown)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.Status(http.StatusOK)
}

func ListNodesByRoot(c *gin.Context) {
	nodes, err := db.ListNodesByRoot(0)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, nodes)
}
