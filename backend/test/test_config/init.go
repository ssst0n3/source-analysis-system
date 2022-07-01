package test_config

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/source-analysis-system/config"
	"os"
)

func Init() {
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDriverName, "mysql"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbHost, "127.0.0.1"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbPort, "16036"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbName, "source-analysis-system"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbUser, "source-analysis-system"))
	awesome_error.CheckFatal(os.Setenv(lightweight_api.EnvDbPasswordFile, "/tmp/secret/MYSQL_PASSWORD_FILE"))
	awesome_error.CheckFatal(os.Setenv(consts.EnvDirSecret, "/tmp/secret"))
	config.LocalListenPort = "16080"
	config.AllowOrigins = append(config.AllowOrigins, "http://127.0.0.1:8080")
}
