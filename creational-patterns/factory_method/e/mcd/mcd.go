package mcd

import (
	b "golang-design-pattern/creational-patterns/factory_method/e/burger"
	m "golang-design-pattern/creational-patterns/factory_method/e/meal"
)

type McSpicy struct {
	b.Burger
}

type AyamPanas struct {
	m.Meal
}

type Kitchen struct{}

func (m *McSpicy) MakeBurger() b.BurgerProduct {
	burger := &McSpicy{
		Burger: b.Burger{
			"bread",
			"spicy chicken",
			"american cheese",
			"mayonnaise",
		},
	}
	return burger
}

func (a *AyamPanas) MakeMeal() m.MealProduct {
	meal := &AyamPanas{
		Meal: m.Meal{
			"spicy",
			"coca cola",
			"fried fries",
		},
	}

	return meal
}

func (k *Kitchen) CreateBurger() b.BurgerProduct {
	return new(McSpicy).MakeBurger()
}

func (k *Kitchen) CreateMeal() m.MealProduct {
	return new(AyamPanas).MakeMeal()
}
