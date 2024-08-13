package main

import "time"

func main() {
	go runSimpleReader()
	go runSimpleWriter()
	time.Sleep(10 * time.Second)
}

var sharedInt = 0
var unUsedInt = 0

func runSimpleReader() {
	for {
		var val int = sharedInt

		if val%10 == 0 {
			unUsedInt = unUsedInt + 1
		}

	}
}

func runSimpleWriter() {
	for {
		sharedInt++
	}
}
