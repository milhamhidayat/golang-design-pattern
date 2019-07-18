package kfc

import (
	b "golang-design-pattern/creational-patterns/factory_method/e/burger"
	m "golang-design-pattern/creational-patterns/factory_method/e/meal"
)

type ZuperKrunch struct {
	b.Burger
}
type SuperBesar struct {
	m.Meal
}

type KFCKitchen struct{}

func (z *ZuperKrunch) MakeBurger() b.BurgerProduct {
	burger := &ZuperKrunch{
		Burger: b.Burger{
			"bread",
			"chicken crispy",
			"mozarella",
			"salad",
		},
	}

	return burger
}

func (s *SuperBesar) MakeMeal() m.MealProduct {
	meal := &SuperBesar{
		Meal: m.Meal{
			"crispy",
			"fanta",
			"mash potato",
		},
	}

	return meal
}

func (k *KFCKitchen) CreateBurger() b.BurgerProduct {
	return new(ZuperKrunch).MakeBurger()
}

func (k *KFCKitchen) CreateMeal() m.MealProduct {
	return new(SuperBesar).MakeMeal()
}
