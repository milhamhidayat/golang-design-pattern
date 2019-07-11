package main

import "fmt"

// this is product
type User struct {
	Name      string
	Role      string
	MinSalary int
	MaxSalary int
}

// director : perform all required step inside builder to create a product
type UserBuilder struct {
	User
}

func (ub *UserBuilder) Build() User {
	return ub.User
}

func (ub *UserBuilder) Name(name string) *UserBuilder {
	ub.User.Name = name
	return ub
}

func (ub *UserBuilder) Role(role string) *UserBuilder {
	if role == "manager" {
		ub.User.MinSalary = 2000
		ub.User.MaxSalary = 6000
	}
	ub.User.Role = role
	return ub
}

func main() {
	ub := &UserBuilder{}
	user := ub.Name("Michael").Role("Staff").Build()
	fmt.Println("++++++++++++++")
	fmt.Printf("%+v\n", user)
	fmt.Println("++++++++++++++")
}
