package main

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/source-analysis-system/config"
	"github.com/ssst0n3/source-analysis-system/db"
	"github.com/ssst0n3/source-analysis-system/router"
)

func main() {
	db.Init()
	r := router.InitRouter()
	awesome_error.CheckFatal(r.Run(fmt.Sprintf(":%s", config.LocalListenPort)))
}
