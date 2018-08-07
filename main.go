package main

import (
	"log"
	"os"

	"github.com/Gigamons/common/logger"

	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/gigamons_api/constants"
	"github.com/Gigamons/gigamons_api/glob"
	"github.com/Gigamons/gigamons_api/router"
	"github.com/Gigamons/gigamons_api/server"
)

var c constants.Config

func init() {
	c := constants.Config{MySQL: consts.MySQLConf{Database: "gigamons", Hostname: "localhost", Username: "root", Port: 3306}, Server: constants.ServerSettings{Port: 23580}}
	if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
		helpers.CreateConfig("config", c)
		logger.Infoln("I've just created a config for you! Please edit config.yml")
		os.Exit(0)
	}
}

func main() {
	helpers.GetConfig("config", &c)
	glob.APIKEY = c.Server.APIKey
	helpers.Connect(&c.MySQL)
	if err := helpers.DB.Ping(); err != nil {
		log.Fatal(err)
	}
	router.Route()
	server.StartServer(c.Server.Hostname, c.Server.Port)
}
