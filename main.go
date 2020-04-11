package main

import (
	"fmt"
	"time"
)

func sendValue(v string, c chan string) {
	fmt.Println("sending in 1 seconds...")
	time.Sleep(1 * time.Second)
	c <- v
	fmt.Println(fmt.Sprintf("sent: '%s'", v))
}

func main() {

	values := make(chan string, 2)
	defer close(values)

	go sendValue("wubba", values)
	time.Sleep(100 * time.Millisecond)
	go sendValue("lubba", values)
	time.Sleep(100 * time.Millisecond)
	go sendValue("dub dub", values)

	fmt.Println("waiting 5 seconds before receiving...")
	time.Sleep(5 * time.Second)

	value := <-values
	fmt.Println(fmt.Sprintf("received: %s", value))

	fmt.Println("waiting 5 seconds before ending...")
	time.Sleep(5 * time.Second)
}
