package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	println("Waiting SIGCINT (Ctrl+C)")
	<-signals
	println("\nSIGINT arrived")
}
