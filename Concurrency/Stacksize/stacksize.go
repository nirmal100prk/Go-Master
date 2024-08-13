package main

import (
	"fmt"
	"runtime"
)

/*
Each goroutine starts with a small stack, typically around 2 kilobytes (KB).
This is much smaller than the default stack size of threads in many operating systems,

	which is often 1 megabyte (MB) or more.

	 This allows a large number of goroutines to be
	 created and run concurrently without consuming excessive memory.
*/
func main() {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Initial memory:%d bytes", m.Alloc)

	go recursiveCall(2)

	runtime.ReadMemStats(&m)
	fmt.Printf("Memory after creating goroutines: %d bytes", m.Alloc)
}

func recursiveCall(n int) {
	if n == 0 {
		return
	}
	recursiveCall(n - 1)
}
