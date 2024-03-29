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

func ListNodesByRoot(c *gin.Context) {
	id, err := ParseId(c)
	if err != nil {
		return
	}
	nodes, err := db.ListNodesByRoot(uint(id), true)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, nodes)
}
