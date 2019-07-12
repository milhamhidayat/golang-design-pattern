package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := &ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	assert.Equal(t, car.Wheels, 4)
	assert.Equal(t, car.Structure, "Car")
	assert.Equal(t, car.Seats, 5)
}
