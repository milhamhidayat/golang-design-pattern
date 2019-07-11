package main

import "fmt"

type MealProduct struct {
	drink string
	food  string
}

type MealBuilder interface {
	SetDrink() MealBuilder
	SetFood() MealBuilder
	GetMeal() MealProduct
}

type MealSpecial1 struct {
	meal MealProduct
}

func (m *MealSpecial1) SetDrink() MealBuilder {
	m.meal.drink = "Fanta"
	return m
}

func (m *MealSpecial1) SetFood() MealBuilder {
	m.meal.food = "McChicken"
	return m
}

func (m *MealSpecial1) GetMeal() MealProduct {
	return m.meal
}

type MealDirector struct {
	builder MealBuilder
}

func (m *MealDirector) Construct(b MealBuilder) {
	b.SetDrink()
	b.SetFood()
	m.builder = b

	// m.builder = b
	// m.builder.SetDrink()
	// m.builder.SetFood()
}

func main() {
	meal1 := &MealSpecial1{}
	build := &MealDirector{}

	build.Construct(meal1)
	fmt.Println(build.builder.GetMeal())
}
