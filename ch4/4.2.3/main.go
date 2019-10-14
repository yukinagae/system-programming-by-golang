package main

import "math"

func primeNumer() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primeNumer()
	for n := range pn {
		println(n)
	}
}
