package main

import "fmt"

func main2() {
	//creat two channels
	intChan := make(chan int)
	stringChan := make(chan string)
	go func() {
		intChan <- 1
	}()
	go func() {
		stringChan <- "string"
	}()
	select {
	case i := <-intChan:
		fmt.Printf("i=%d", i)
	case s := <-stringChan:
		fmt.Printf("s=%s", s)
	}
}
