package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	// //oncefunc
// 	// once := sync.OnceFunc(Helloworld)
// 	// for i := 0; i < 5; i++ {
// 	// 	once()
// 	// }

// 	//onceValue -subsequent invocations yield the same value.

// 	randNum := func() int {

// 		return rand.IntN(100) % 10
// 	}
// 	getNum := sync.OnceValue(randNum)
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(getNum())
// 	}
// }

// singleton pattern
type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

//oncefunc

func Helloworld() {
	fmt.Println("Hello, world!")
}

//oncevalue
