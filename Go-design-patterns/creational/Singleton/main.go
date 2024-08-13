package main

import "sync"

func main() {

}

/*
The Singleton pattern ensures that a class has only one instance
 and provides a global point of access to that instance.
*/

type Logger struct {
	Logs []string
}

var instance *Logger
var once sync.Once

func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{}
		instance.Logs = make([]string, 0)

	})
	return instance
}

func (l *Logger) Log(message string) {
	l.Logs = append(l.Logs, message)
}
