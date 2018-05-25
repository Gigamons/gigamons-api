package constants

import (
	"github.com/Gigamons/common/consts"
)

type ServerSettings struct {
	APIKey   string
	Port     int16
	Hostname string
}

type Config struct {
	Server ServerSettings
	MySQL  consts.MySQLConf
}
