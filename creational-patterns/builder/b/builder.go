package main

import (
	"errors"
	"fmt"
)

type User struct {
	name      string
	role      string
	minSalary int
	maxSalary int
}

type UserOption func(user *User) error

func BuildUser(opts ...UserOption) (*User, error) {
	var user User
	for _, opt := range opts {
		err := opt(&user)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func WithName(name string) UserOption {
	return func(user *User) error {
		user.name = name
		return nil
	}
}

func WithRole(role string) UserOption {
	return func(user *User) error {
		if role != "manager" && role != "sales" {
			return errors.New("invalid role")
		}
		if role == "manager" {
			user.minSalary = 40000
			user.maxSalary = 60000
		}
		user.role = role
		return nil
	}
}

func main() {
	user, err := BuildUser(
		WithName("Michael"),
		WithRole("sales"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("++++++++++++++")
	fmt.Printf("%+v\n", user)
	fmt.Println("++++++++++++++")
}
