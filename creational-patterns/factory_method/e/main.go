package main

import (
	"fmt"
	r "golang-design-pattern/creational-patterns/factory_method/e/restaurant"
)

func main() {
	restaurant := &r.RestaurantFood{}

	burger, err := restaurant.GetBurger(1)
	if err != nil {
		panic(err)
	}

	meal, err := restaurant.GetMeal(1)

	if err != nil {
		panic(err)
	}

	fmt.Println("++++++++++++++")
	fmt.Printf("%+v\n", burger)
	fmt.Println("++++++++++++++")

	fmt.Println("==============")
	fmt.Printf("%+v\n", meal)
	fmt.Println("==============")
}
