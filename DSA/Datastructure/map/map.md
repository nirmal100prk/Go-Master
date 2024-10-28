**MAP**

https://go.dev/blog/maps

# why maps are not concurrency safe ?
Maps in Go are not concurrent-safe because they don’t provide built-in synchronization, which can lead to race conditions when accessed by multiple goroutines.
Race conditions in maps can cause unexpected behavior, data corruption, or runtime panics.
To make maps concurrent-safe, use sync.Mutex, sync.RWMutex, or sync.Map, depending on your specific needs.
sync.Map is a built-in concurrent map implementation that simplifies concurrency management but might not always offer the best performance compared to custom locking strategies.

