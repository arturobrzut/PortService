package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"port/pkg/db"
	"port/pkg/server"
	"port/pkg/signal"
)

func main() {
	log.Info("Start Port service")

	dbHandler, err := db.NewDbHandler()
	if err != nil {
		log.Fatal("Cannot set database handler. ", err)
		os.Exit(1)
	}

	config, err := server.SetupService(dbHandler)
	if err != nil {
		log.Fatal("Cannot set service config. ", err)
		os.Exit(2)
	}

	signalCh := make(chan bool)
	serviceCh := make(chan bool)
	repeat := true

	go signal.Handler(signalCh, dbHandler)
	go func(config *server.ServiceConfig) {
		err := server.Start(config)
		if err != nil {
			log.Error("Error during starting service. ", err)
		}
		serviceCh <- false
	}(config)

	for repeat {
		select {
		case repeat = <-signalCh:
			if repeat == false {
				log.Info("Stop service by signal")
			}
		case <-serviceCh:
			{
				log.Info("Stop service")
				repeat = false
			}
		}
	}
}
