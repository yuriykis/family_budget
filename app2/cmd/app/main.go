package main

import (
	"app/internal/app/apiserver"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const defaultConfigPath = "configs/server.yml"

func main() {
	config := apiserver.NewConfig()
	_, err := yaml.Marshal(defaultConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
