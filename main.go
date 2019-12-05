package main

import (
	"flag"
	"iris/config"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
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

	//init cache
	if err := config.InitCache(config.GetCache()); err != nil {
		log.Fatalf("Fatal Error: can't initialize cache!!!\n%s", err)
	}

	// init db clients
	if err := config.InitDB(config.GetDB()); err != nil {
		log.Fatalf("Fatal Error: can't initialize mysql!!!\n%s", err)
	}

	config.NewApp()

	// waite for exit signal
	exit := make(chan os.Signal)
	stopSignal := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	signal.Notify(exit, stopSignal...)

	// catch exit signal
	sign := <-exit
	config.LogInfo("stop by exit signal '%s'", sign)
}
