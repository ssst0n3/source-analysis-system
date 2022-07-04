package node_relation

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/response"
	"github.com/ssst0n3/source-analysis-system/api/v1/node"
	"github.com/ssst0n3/source-analysis-system/db"
	"github.com/ssst0n3/source-analysis-system/model"
	"net/http"
)

const ResourceName = "node_relation"

var Resource = lightweight_api.NewResource(ResourceName, model.SchemaNodeRelation.Table, model.NodeRelation{}, "")

func ListByRoot(c *gin.Context) {
	id, err := node.ParseId(c)
	if err != nil {
		return
	}
	nodeRelations, err := db.ListNodeRelationsByRootWithRoot(uint(id))
	if err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	c.JSON(http.StatusOK, nodeRelations)
}

func UpdateNodeRelationByNode(c *gin.Context) {
	id, err := node.ParseId(c)
	if err != nil {
		return
	}
	err = node.Resource.MustResourceExistsById(c, uint(id))
	if err != nil {
		return
	}
	var nodeRelation model.NodeRelation
	if err := c.ShouldBindJSON(&nodeRelation); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	err = db.DB.Debug().Model(model.NodeRelation{}).Where(model.SchemaNodeRelation.FieldsByName["Node"].DBName, id).Updates(
		&nodeRelation,
	).Error
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else {
		response.UpdateSuccess200(c)
		return
	}
}
