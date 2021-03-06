package main

import "fmt"

// produk
// object terkahir yang diciptkan oleh builder
// menentukan elemen penyusunnya, blueprint
type PizzaProduct struct {
	dough   string
	sauce   string
	topping string
}

// builder
// interface untuk membuat sebuah produk
type PizzaBuilder interface {
	SetDough() PizzaBuilder
	SetSauce() PizzaBuilder
	SetTopping() PizzaBuilder
	GetPizza() PizzaProduct
}

// concrete builder
// menyediakan fungsi untuk membuat sebuah produk, mengimplementasikan interface dari builder
type MeatPizza struct {
	ingredients PizzaProduct
}

func (m *MeatPizza) SetDough() PizzaBuilder {
	m.ingredients.dough = "Large"
	return m
}

func (m *MeatPizza) SetSauce() PizzaBuilder {
	m.ingredients.sauce = "Tomato"
	return m
}

func (m *MeatPizza) SetTopping() PizzaBuilder {
	m.ingredients.topping = "Wagyu meat"
	return m
}

func (m *MeatPizza) GetPizza() PizzaProduct {
	return m.ingredients
}

// director
// mendefiniskan urutan dalam menyusun produk
// menggunakan interface yang diimplementasikan oleh concrete builder
type PizzaDirector struct {
	builder PizzaBuilder
}

// menerima argumen type yand mengimplementasi PizzaBuilder interface
func (p *PizzaDirector) SetPizza(b PizzaBuilder) {
	p.builder = b
}

func (p *PizzaDirector) Construct() {
	p.builder.SetDough().SetSauce().SetTopping()
}

func main() {
	pizza1 := &MeatPizza{}

	build := &PizzaDirector{}
	build.SetPizza(pizza1)

	build.Construct()
	fmt.Println(build.builder.GetPizza())
}
