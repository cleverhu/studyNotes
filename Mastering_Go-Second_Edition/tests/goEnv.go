package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			for true {
				<-time.After(1 * time.Second)
			}
		}()
	}

	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
}
