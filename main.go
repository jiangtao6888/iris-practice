package main

import (
	"flag"
	"iris/config"
	"log"
	"runtime"
)

var configFile = flag.String("c", "config.toml", "config file")

func main() {
	// parse flag
	flag.Parse()
	// set max cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := config.InitConfig(*configFile); err != nil {
		log.Fatalf("Fatal Error: can't parse config file!!!\n%s", err)
	}

	// init log
	if err := config.InitLogger(config.GetLog()); err != nil {
		log.Fatalf("Fatal Error: can't initialize logger!!!\n%s", err)
	}


	//Log().WithFields(logrus.Fields{
	//	"animal": "walrus",
	//	"size":   10,
	//}).Info("A group of walrus emerges from the ocean")

	config.NewApp()
}
