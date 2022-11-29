package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/snail007/goproxy/services"
)

const APP_VERSION = "3.0-c18s-1.0"

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("err : %s", err)
	}
	Clean(&service.S)
}
func Clean(s *services.Service) {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for range signalChan {
			fmt.Println("\nReceived an interrupt, stopping services...")
			(*s).Clean()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
