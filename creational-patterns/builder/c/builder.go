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
	SetDough()
	SetSauce()
	SetTopping()
	GetPizza()
}

// concrete builder
// menyediakan fungsi untuk membuat sebuah produk, mengimplementasikan interface dari builder
type MeatPizza struct {
	ingredients PizzaProduct
}

func (m *MeatPizza) SetDough() {
	m.ingredients.dough = "Large"
}

func (m *MeatPizza) SetSauce() {
	m.ingredients.sauce = "Tomato"
}

func (m *MeatPizza) SetTopping() {
	m.ingredients.topping = "Wagyu meat"
}

func (m *MeatPizza) GetPizza() {
	fmt.Println("++++++++++++++")
	fmt.Printf("%+v\n", m.ingredients)
	fmt.Println("++++++++++++++")
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
	p.builder.SetDough()
	p.builder.SetSauce()
	p.builder.SetTopping()
}

func main() {
	pizza1 := &MeatPizza{}

	build := &PizzaDirector{}
	build.SetPizza(pizza1)

	build.Construct()
	build.builder.GetPizza()
}
