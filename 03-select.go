package main

import (
	"fmt"
	"time"
)

// write
func write(c1 chan string) {
	for {
		select {
		//write data
		case c1 <- "hello":
			fmt.Println(c1)
		default:
			fmt.Println("channel is full")
		}
		time.Sleep(time.Second * 2)
	}
}
func main() {
	//creat channel
	output1 := make(chan string, 10)
	go write(output1)
	//fetch data
	for s := range output1 {
		fmt.Println(s)
		time.Sleep(time.Second * 1)
	}

}
