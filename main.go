package main

import (
	"flag"
	"log"

	"redis_golang/config"
	"redis_golang/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the redis server")
	flag.IntVar(&config.Port, "port", 6379, "port for the redis server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("starting the redis server..")
	server.RunAsyncTCPServer()
}