package main

import "fmt"

/*
Proxy is a structural design pattern that provides an object that acts
 as a substitute for a real service object used by a client.
 A proxy receives client requests, does some work (access control, caching, etc.)
 and then passes the request to a service object.

 The proxy should use the same concrete struct. It should use the concrete
  struct (embed it) and perform an operation. such that this operations is performed
  before the concrete method.

*/

type IDrive interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	age int
}

type CarProxy struct {
	Cr Car
	Dr *Driver
}

func (c *CarProxy) Drive() {
	if c.Dr.age > 18 {
		c.Cr.Drive()
	} else {
		fmt.Println("Driver is too young to drive")
	}
}

func NewCarProxy(age int) IDrive {
	return &CarProxy{
		Cr: Car{},
		Dr: &Driver{age: age},
	}
}

func main() {
	Ncp := NewCarProxy(19)
	Ncp.Drive()
}
