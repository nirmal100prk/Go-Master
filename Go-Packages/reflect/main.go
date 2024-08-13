package main

import (
	"fmt"
	"reflect"
)

func main() {
	d := Details{
		fname:   "abc",
		lname:   "efg",
		age:     10,
		balance: 1.2,
	}
	showType(d)

}

type Details struct {
	fname   string
	lname   string
	age     int32
	balance float32
}

func showType(i interface{}) {
	t := reflect.TypeOf(i)
	k := t.Kind()

	fmt.Println("Type of first interface :", t)
	fmt.Println("Kind of first interface: ", k)
	fmt.Println(reflect.ValueOf(i))
	fmt.Println(reflect.ValueOf(i).Kind())
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		value := reflect.ValueOf(i)
		numberOfField := value.NumField()

		for i := 0; i < numberOfField; i++ {
			fmt.Println(value.Field(i), value.Field(i).Kind())
		}

	}

}
