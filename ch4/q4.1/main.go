package main

import (
	"fmt"
	"time"
)

func main() {

	println("timer start")

	fmt.Printf("%v\n", time.Now())

	timer := <-time.After(5 * time.Second)

	fmt.Printf("%v\n", timer)
	println("timer end")
}
