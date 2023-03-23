package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"transactionservice/internal/app/apiserver"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var configPath = "../../configs/server.yml"

func init() {
	flag.StringVar(&configPath, "config-path", configPath, "path to config file")
}

func main() {
	config := apiserver.NewConfig()
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	configPath := filepath.Join(wd, configPath)
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
