package config

import (
	"os"
	"strings"
)

const (
	EnvLocalListenPort = "LOCAL_LISTEN_PORT"
	EnvAllowOrigins    = "ALLOW_ORIGINS"
)

var (
	LocalListenPort = os.Getenv(EnvLocalListenPort)
	AllowOrigins    []string
)

func init() {
	origins := os.Getenv(EnvAllowOrigins)
	if len(strings.TrimSpace(origins)) == 0 {
		AllowOrigins = nil
	} else {
		AllowOrigins = strings.Split(origins, ",")
	}
}
