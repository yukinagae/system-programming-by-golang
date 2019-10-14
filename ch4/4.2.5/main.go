package main

import "context"

func main() {
	println("start sub()")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		println("sub() is finished")
		cancel()
	}()

	<-ctx.Done()
	println("all tasks are finished")
}
