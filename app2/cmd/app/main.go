package main

import (
	"app/internal/app/apiserver"
	"flag"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var configPath = "configs/server.yml"

func init() {
	flag.StringVar(&configPath, "config-path", configPath, "path to config file")
}

func main() {
	config := apiserver.NewConfig()
	v, err := yaml.Marshal(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(v, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
