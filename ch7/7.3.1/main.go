package main

import (
	"net"
	"time"
)

const interval = 10 * time.Second

func main() {
	println("Start tick server at 224.0.0.1:9999")
	conn, err := net.Dial("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	start := time.Now()

	wait := start.Truncate(interval).Add(interval).Sub(start)
	time.Sleep(wait)
	ticker := time.Tick(interval)

	for now := range ticker {
		conn.Write([]byte(now.String()))
		println("Tick: " + now.String())
	}
}
