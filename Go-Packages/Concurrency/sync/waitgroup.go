// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go increment(&wg)
// 	}
// 	wg.Wait()
// 	time.Sleep(5 * time.Second)
// 	fmt.Printf("counter: %v", counter)
// }

// var counter int
// var mutex sync.Mutex

// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(5 * time.Second)
// 	for i := 0; i < 1000; i++ {

// 		counter++

// 	}

// }
package main

var count int

func race() {
	count++
}

// func main() {
// 	go race()
// 	go race()
// 	time.Sleep(1 * time.Second)
// }
