package main

import (
	"categoryservice/internal/app/apiserver"
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

var configPath = "../../configs/server.yml"

func init() {
	flag.StringVar(&configPath, "config-path", configPath, "path to config file")
}

func main() {
	config := apiserver.NewConfig()
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// configPath := filepath.Join(wd, configPath)
	// flag.Parse()
	// configFile, err := ioutil.ReadFile(configPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = yaml.Unmarshal(configFile, config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	config.BindAddr = os.Getenv("BIND_ADDR")
	config.LogLevel = os.Getenv("LOG_LEVEL")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	config.DatabaseURL = "mongodb://" + dbHost + ":" + dbPort

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
