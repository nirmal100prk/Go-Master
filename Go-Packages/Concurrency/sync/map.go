package main

import (
	"fmt"
	"sync"
)

// for managing  shared datastructures dafely and efficiently
// concurrent safe map implementation

// introduced in GO 1.9

/*
Key Features of sync.Map
Concurrent Access: sync.Map allows safe concurrent access to the map from multiple
 goroutines without the need for external synchronization mechanisms like mutexes.
Dynamic Key-Value Pairs: Keys and values in sync.Map can be dynamically added,
updated, or deleted concurrently.
Lazy Initialization: sync.Map does not require initialization and can be used
 immediately after declaration.
No Copying: Unlike traditional maps, sync.Map does not perform copies during
 read or write operations, which can improve performance in certain scenarios.
*/

func main() {
	var m sync.Map

	m.Store("key1", "value1")
	m.Store("key2", "value2")

	m.Delete("key1")
	val, ok := m.Load("key2")
	if ok {
		fmt.Println(val.(string))
	}

	m.Store("key3", "value3")

	/// range over sync map
	m.Range(func(key, val any) bool {
		fmt.Printf("key: %v, value: %v\n", key, val)
		return true
	})

	// concurrent access

	var mp sync.Map
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mp.Store(n, n*10)

		}(i)
	}
	//fmt.Println(mp)
	wg.Wait()
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			val, ok := mp.Load(n)
			if ok {
				fmt.Println(val.(int))
			}

		}(i)
	}
	wg.Wait()
}
