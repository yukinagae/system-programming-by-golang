package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {

	clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	os.Remove(clientPath)

	conn, err := net.ListenPacket("unixgram", clientPath)
	if err != nil {
		panic(err)
	}

	serverPath := filepath.Join(os.TempDir(), "unixdomainsocket-server")
	unixServerAddr, err := net.ResolveUnixAddr(
		"unixgram",
		serverPath,
	)
	var serverAddr net.Addr = unixServerAddr
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	log.Println("Sending to Server")

	_, err = conn.WriteTo([]byte("Hello from Client"), serverAddr)
	if err != nil {
		panic(err)
	}

	log.Println("Receiving from Server")
	buffer := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}
	log.Printf("Received: %s\n", string(buffer[:length]))
}
