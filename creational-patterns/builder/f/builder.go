package main

import "fmt"

// product
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// builder
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// director
type ManufacturingDirector struct{}

func (f *ManufacturingDirector) Construct() {

}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {

}

// concrete builder
type CarBuilder struct{}

func (c *CarBuilder) SetWheels() BuildProcess {
	return nil
}

func (c *CarBuilder) SetSeats() BuildProcess {
	return nil
}

func (c *CarBuilder) SetStructure() BuildProcess {
	return nil
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return VehicleProduct{}
}

// concrete builder
type BikeBuilder struct{}

func (b *BikeBuilder) SetWheels() BuildProcess {
	return nil
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	return nil
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	return nil
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return VehicleProduct{}
}

func main() {
	fmt.Println("this is dummy driver")
}
