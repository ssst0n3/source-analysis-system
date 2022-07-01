package node_relation

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/source-analysis-system/db"
	"github.com/ssst0n3/source-analysis-system/model"
	"net/http"
	"strconv"
)

const ResourceName = "node_relation"

var Resource = lightweight_api.NewResource(ResourceName, model.SchemaNodeRelation.Table, model.NodeRelation{}, "")

func ListByRoot(c *gin.Context) {
	rootId := c.Param("id")
	id, err := strconv.ParseUint(rootId, 10, 64)
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	nodeRelations, err := db.ListNodeRelationsByRootWithRoot(uint(id))
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	c.JSON(http.StatusOK, nodeRelations)
}
