package restaurant

import (
	"errors"
	b "golang-design-pattern/creational-patterns/factory_method/e/burger"
	kfc "golang-design-pattern/creational-patterns/factory_method/e/kfc"
	mcd "golang-design-pattern/creational-patterns/factory_method/e/mcd"
	m "golang-design-pattern/creational-patterns/factory_method/e/meal"
)

type Restaurant interface {
	CreateBurger() b.BurgerProduct
	CreateMeal() m.MealProduct
}

type RestaurantFood struct{}

func (r *RestaurantFood) GetBurger(num int) (b.BurgerProduct, error) {
	if num == 1 {
		return new(kfc.KFCKitchen).CreateBurger(), nil
	} else if num == 2 {
		return new(mcd.McDKitchen).CreateBurger(), nil
	} else {
		return nil, errors.New("option is not valid")
	}
}

func (r *RestaurantFood) GetMeal(num int) (m.MealProduct, error) {
	if num == 1 {
		return new(kfc.KFCKitchen).CreateMeal(), nil
	} else if num == 2 {
		return new(mcd.McDKitchen).CreateMeal(), nil
	} else {
		return nil, errors.New("option is not valid")
	}
}
