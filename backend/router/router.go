package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	v1 "github.com/ssst0n3/source-analysis-system/api/v1"
	"github.com/ssst0n3/source-analysis-system/config"
	"net/http"
)

type ResponsePing struct {
	Message string `json:"message" example:"pong"`
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, ResponsePing{Message: "pong"})
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	// cors
	if config.AllowOrigins != nil {
		corsConfig := cors.DefaultConfig()
		//corsConfig.AllowAllOrigins = true
		corsConfig.AllowCredentials = true
		corsConfig.AllowOrigins = append(corsConfig.AllowOrigins, config.AllowOrigins...)
		router.Use(cors.New(corsConfig))
	}
	//frontend
	{
		router.Use(static.Serve("/", static.LocalFile("./html", false)))
	}
	// ping pong
	{
		router.GET("/ping", Ping)
	}
	{
		v1.InitRouter(router)
	}
	return router
}
