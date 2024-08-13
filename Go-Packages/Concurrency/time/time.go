package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick at", time.Now())
		case <-done:
			ticker.Stop()
			fmt.Println("Done")
			return
		}

	}

}
