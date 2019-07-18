package burger

type Burger struct {
	Bun    string
	Meat   string
	Cheese string
	Option string
}

type BurgerProduct interface {
	MakeBurger() BurgerProduct
}
