package main

import "sync"

/*
pool is thread safe
pool is a place where you can keep temporary objects for later reuse

When you’ve got a lot of goroutines running at once, they often need similar objects. Imagine running go f() multiple times concurrently.

If each goroutine creates its own objects, memory usage can quickly increase and this puts a strain on the garbage collector because it has
to clean up all those objects once they’re no longer needed


*/

var pool sync.Pool = sync.Pool{
	New: func() any {
		return new(int)
	},
}
