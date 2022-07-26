package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/source-analysis-system/api/v1/node"
	"github.com/ssst0n3/source-analysis-system/api/v1/node_relation"
)

func InitRouter(router *gin.Engine) {
	nodeRelationGroup := router.Group(node_relation.Resource.BaseRelativePath)
	{
		nodeRelationGroup.GET("/:id", node_relation.ListByRoot)
		nodeRelationGroup.POST("", node_relation.Resource.CreateResource)
		nodeRelationGroup.PUT("/node/:id", node_relation.UpdateNodeRelationByNode)
	}
	nodeGroup := router.Group(node.Resource.BaseRelativePath)
	{
		nodeGroup.GET("/list/:id", node.ListNodesByRoot)
		nodeGroup.GET("/matrix/:id", node.Matrix)
		nodeGroup.PUT("/:id", node.UpdateMarkdown)
		nodeGroup.POST("", node.Resource.CreateResource)
	}
}
