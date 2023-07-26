package main

import (
	"fmt"
	"time"
)

func test1(c1 chan string) {
	time.Sleep(5 * time.Second)
	c1 <- "string1"

}
func test2(c2 chan string) {
	time.Sleep(2 * time.Second)
	c2 <- "string2"

}
func main1() {
	//create channels
	output1 := make(chan string)
	output2 := make(chan string)
	go test1(output1)
	go test2(output2)

	select {
	case s1 := <-output1:
		fmt.Printf("s1=%v", s1)
	case s2 := <-output2:
		fmt.Printf("s2=%v", s2)
	}
}
