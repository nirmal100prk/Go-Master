package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

func worker(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d: %d\n", id, i)
		time.Sleep(time.Millisecond * 500) // Simulates work
	}
}

func main() {
	// even though 2 threads are set, more threads will be spawn by go runtime
	// for i/o waits, GC, or syscalls in background
	runtime.GOMAXPROCS(2) // Limit OS threads to 2
	// Create a CPU profile file
	f, err := os.Create("profile.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	traceFile, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer traceFile.Close()

	err = trace.Start(traceFile)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		go worker(i) // Spawn multiple goroutines
	}

	time.Sleep(3 * time.Second) // Wait for goroutines to finish
}
