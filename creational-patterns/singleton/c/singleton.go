package main

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}

func main() {
	s := New()
	s["this"] = "that"

	fmt.Println("++++++++++++++")
	fmt.Printf("%+v\n", s)
	fmt.Println("++++++++++++++")

	s2 := New()
	fmt.Println("==============")
	fmt.Printf("%+v\n", s2)
	fmt.Println("==============")
}
