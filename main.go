package main

import (
	"os"
	"os/signal"
	"ping/tools"
	"syscall"
)

func main() {
	ping := tools.NewPingTools("8.8.8.8")
	ping.Execute()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
