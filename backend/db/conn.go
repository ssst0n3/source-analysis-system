package db

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	awesome_error.CheckFatal(lightweight_api.InitGormDB())
	DB = lightweight_api.DB
	awesome_error.CheckFatal(Migrate())
}
