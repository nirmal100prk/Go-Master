package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	time.Sleep(5 * time.Second)

}

func simpleChannel() {
	t := time.Now()

	s := make(chan int, 1)
	s <- 1
	go func() {

		s <- 2
	}()

	fmt.Println(time.Since(t))
}

// already filled unbuffered channel
func DeadLockSample1() {
	c := make(chan int)
	c <- 1
	go func() {
		c <- 2
	}()
	fmt.Println(<-c)
}

// reading from an empty channel
func DeadLockSample2() {
	c := make(chan int, 1)
	go func() {}()
	<-c
}

// without receiver
func DeadLockSample3() {
	c := make(chan int)
	go func() {
		c <- 1
	}()

}
