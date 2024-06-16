package main

import (
	"log"
	"src/config"
	"src/internal/app"
	"src/pkg/db"
)

func main() {
	conf, err := config.LoadNewConfig()
	if err != nil {
		log.Fatal(err)
	}
	appConfig, err := conf.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}
	poolConnection, err := db.NewPsqlPoolConnection(appConfig)
	if err != nil {
		log.Fatal("Error with connect DB:", err)
	}
	defer poolConnection.Close()
	app.NewApp(poolConnection, &appConfig).Run()
}
