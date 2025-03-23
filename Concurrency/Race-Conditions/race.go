package main

import (
	"fmt"
	"sync"
	"time"
)

// step 1 : run without locking, to see race conditions
// step 2 : only lock write, still a race condition will be present on read
// step 3 : lock read and write, no race condition will be present

// go run -race race.go
func main() {
	Race()
	// adding mutex in this print statement, even though it is ran after Race func is called.
	// addding mutex to synchronize  with the writes
	// Go's race detector flags any unsynchronized access (read or write) to
	// a shared variable when at least one of those accesses is a write.
	//  Since the read in fmt.Println is unsynchronized,
	// it triggers a data race warning.
	m.Lock()
	fmt.Println("Counter:", counter)
	m.Unlock()
}

var counter int

var m sync.Mutex

func Race() {
	for i := 0; i < 5; i++ {
		go func() {
			m.Lock()
			counter++
			m.Unlock()
		}()
	}
	time.Sleep(2 * time.Second)

}

// var sharedInt = 0
// var unUsedInt = 0

// func runSimpleReader() {
// 	for {
// 		var val int = sharedInt

// 		if val%10 == 0 {
// 			unUsedInt = unUsedInt + 1
// 		}

// 	}
// }

// func runSimpleWriter() {
// 	for {
// 		sharedInt++
// 	}
// }
