package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()

			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			for {
				// timeout: 5 seconds
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))

				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}

				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				println(string(dump))

				content := "Hello World\n"

				// HTTP/1.1
				// ContentLength
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader(content)),
				}
				response.Write(conn)
			}
		}()
	}
}
