package main

import (
	"app/internal/app/apiserver"
	"flag"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var configPath = "server.yml"

func init() {
	flag.StringVar(&configPath, "config-path", configPath, "path to config file")
}

func main() {
	config := apiserver.NewConfig()
	mode := os.Getenv("GIN_MODE")
	if mode == "debug" {
		configPath = "../../configs/server.yml"
	}
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
