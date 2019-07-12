package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilderPattern(t *testing.T) {

	t.Run("car builder", func(t *testing.T) {
		manufacturingComplex := &ManufacturingDirector{}

		carBuilder := &CarBuilder{}
		manufacturingComplex.SetBuilder(carBuilder)
		manufacturingComplex.Construct()

		car := carBuilder.GetVehicle()

		assert.Equal(t, car.Wheels, 4)
		assert.Equal(t, car.Structure, "Car")
		assert.Equal(t, car.Seats, 5)
	})

	t.Run("bike builder", func(t *testing.T) {
		manufacturingComplex := &ManufacturingDirector{}

		bikeBuilder := &BikeBuilder{}
		manufacturingComplex.SetBuilder(bikeBuilder)
		manufacturingComplex.Construct()

		motorbike := bikeBuilder.GetVehicle()
		motorbike.Seats = 1

		assert.Equal(t, motorbike.Wheels, 2)
		assert.Equal(t, motorbike.Structure, "Motorbike")
	})

}
