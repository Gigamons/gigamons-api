package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
	"github.com/gigamons_api/Router"
	"github.com/gigamons_api/constants"
	"github.com/gigamons_api/glob"
	"github.com/gigamons_api/server"
)

var c constants.Config

func init() {
	c := constants.Config{MySQL: consts.MySQLConf{Database: "gigamons", Hostname: "localhost", Username: "root", Port: 3306}, Server: constants.ServerSettings{Port: 23580}}

	if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
		helpers.CreateConfig("config", c)
		fmt.Println("I've just created a config for you! Please edit config.yml")
		os.Exit(0)
	}
}

func main() {
	helpers.GetConfig("config", &c)
	glob.APIKEY = c.Server.APIKey
	helpers.Connect(c.MySQL)
	if err := helpers.DB.Ping(); err != nil {
		log.Fatal(err)
	}
	Router.Route()
	server.StartServer(c.Server.Hostname, c.Server.Port)
}
