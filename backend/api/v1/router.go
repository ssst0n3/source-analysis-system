package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/source-analysis-system/api/v1/node"
	"github.com/ssst0n3/source-analysis-system/api/v1/node_relation"
)

func InitRouter(router *gin.Engine) {
	nodeRelationGroup := router.Group(node_relation.Resource.BaseRelativePath)
	{
		nodeRelationGroup.GET("/list/:id", node_relation.ListByRoot)
		nodeRelationGroup.GET("/count/:id", node_relation.CountByRoot)
		nodeRelationGroup.POST("", node_relation.Resource.CreateResource)
		nodeRelationGroup.PUT("/node/:id", node_relation.UpdateNodeRelationByNode)
		//nodeRelationGroup.POST("/unlink/:id", node_relation.UnlinkNodeFromParent)
		nodeRelationGroup.DELETE("/hide/:id", node_relation.HideNodeFromParent)
	}
	nodeGroup := router.Group(node.Resource.BaseRelativePath)
	{
		nodeGroup.GET("/list/:id", node.ListNodesByRoot)
		nodeGroup.GET("/matrix/:id", node.Matrix)
		nodeGroup.PUT("/:id", node.Resource.UpdateResource)
		nodeGroup.POST("", node.Resource.CreateResource)
		nodeGroup.DELETE("/:id", node.Resource.DeleteResource)
	}
}
