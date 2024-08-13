package abstractfactory

import "fmt"

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

type CarTyre struct{}

func (c *CarTyre) Type() string {
	return "Im a car tyre"
}

type BikeTyre struct{}

func (b *BikeTyre) Type() string {
	return "Im a bike tyre"
}

//product interface
type Vehicle interface {
	Drive() string
}

//product interface
type Tyre interface {
	Type() string
}

//factory - creator interface
type VehicleFactory interface {
	NewVehicle() Vehicle
	CreateTyre() Tyre
}

//concrete creator
type CarFactory struct{}

func (c *CarFactory) NewVehicle() Tyre {
	return &CarTyre{}
}

func (c *CarFactory) CreateTyre() Vehicle {
	return &Car{}
}

type BikeFactory struct{}

func (c *BikeFactory) NewVehicle() Vehicle {
	return &Bike{}
}

func (c *BikeFactory) CreateTyre() Tyre {
	return &BikeTyre{}
}
