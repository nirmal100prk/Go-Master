package main

import "fmt"

/*
useful when we need to create objects without
specifying their concrete types up front.
*/

//concrete product
type Car struct{}

func (c *Car) Drive() string {
	a := "car"
	return fmt.Sprintf("im driving a %s", a)
}

type Bike struct{}

func (c *Bike) Drive() string {
	return "im driving a bike"
}

//product interface
type Vehicle interface {
	Drive() string
}

//factory - creator interface
type VehicleFactory interface {
	NewVehicle() Vehicle
}

//concrete creator
type CarFactory struct{}

func (c *CarFactory) NewVehicle() Vehicle {
	return &Car{}
}

type BikeFactory struct{}

func (c *BikeFactory) NewVehicle() Vehicle {
	return &Bike{}
}

func GetVehicle(factory VehicleFactory) Vehicle {
	return factory.NewVehicle()
}

func main() {
	v := GetVehicle(&CarFactory{})
	fmt.Println(v.Drive())
}
