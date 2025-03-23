Go is well-suited for Concurrency because of its lightweight Goroutines and built-in Channel type. Goroutines are lightweight threads that can be created easily and have low overhead, allowing for the efficient creation of thousands or even millions of concurrent processes. Channels are built-in data structures that facilitate communication between Goroutines, enabling safe and efficient synchronization of data access.


## memory management 
Avoid excessive use of new() or pointers when not needed
Excessive use of pointers increases GC work due to more heap allocations.
Heap allocations are slower than stack allocations and put pressure on the GC.
Use pointers only when you need shared mutable state or to avoid copying large objects

Go GC is designed to handle lots of small objects efficiently.
Large allocations inside loops force the GC to pause more often due to memory pressure.
Minimizing large heap allocations improves throughput and reduces GC pauses.
Use stack-based data structures when possible, such as slices, maps, and structs.
Use sync.Pool for caching frequently allocated objects to reduce GC pressure.



