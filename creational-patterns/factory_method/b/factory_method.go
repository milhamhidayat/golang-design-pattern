package main

import "fmt"

// 'product' interface
type IFactory interface {
	Drive(miles int)
}

// 'concrete product'
type Scooter struct{}

func (s *Scooter) Drive(miles int) {
	fmt.Printf("drive the scooter : %d km \n", miles)
}

type Bike struct{}

func (b *Bike) Drive(miles int) {
	fmt.Printf("drive the bike : %d km \n", miles)
}

type BMW struct{}

func (b *BMW) Drive(miles int) {
	fmt.Printf("drive the bmw : %d km \n", miles)
}

// creator
type VehicleFactory interface {
	GetVehicle(vehicle string) IFactory
}

// concreate creator
type MotorcycleFactory struct{}

func (c *MotorcycleFactory) GetVehicle(vehicle string) IFactory {
	var vehicleType IFactory
	if vehicle == "a" {
		vehicleType = &Scooter{}
	} else if vehicle == "b" {
		vehicleType = &Bike{}
	}

	return vehicleType
}

type CarFactory struct{}

func (c *CarFactory) GetVehicle(vehicle string) IFactory {
	var vehicleType IFactory
	if vehicle == "c" {
		vehicleType = &BMW{}
	}
	return vehicleType
}

func main() {
	mfactory := &MotorcycleFactory{}

	scooter := mfactory.GetVehicle("a")
	scooter.Drive(10)

	bike := mfactory.GetVehicle("b")
	bike.Drive(30)

	cfactory := &CarFactory{}
	bmw := cfactory.GetVehicle("c")
	bmw.Drive(100000)
}
