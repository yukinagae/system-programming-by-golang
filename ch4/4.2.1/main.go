package main

func main() {
	println("start sub()")

	done := make(chan bool)

	go func() {
		println("sub() is finished")
		done <- true
	}()

	<-done
	println("all tasks are finished")
}
