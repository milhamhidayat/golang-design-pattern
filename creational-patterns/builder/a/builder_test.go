package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserBuilder(t *testing.T) {
	t.Run("Should create empty struct with default value nil", func(t *testing.T) {
		ub := &UserBuilder{}
		want := &User{}
		assert.Equal(t, &ub.User, want)
	})

	t.Run("Should set user name", func(t *testing.T) {
		want := "Tony"
		ub := &UserBuilder{}
		user := ub.Name(want).Build()
		assert.Equal(t, user.Name, want)
	})

	t.Run("Should set role as manager", func(t *testing.T) {
		want := User{
			"Tony", "manager", 2000, 6000,
		}
		ub := &UserBuilder{}
		user := ub.Name("Tony").Role("manager").Build()
		assert.Equal(t, user, want)
	})

	t.Run("Should set role as staff", func(t *testing.T) {
		want := User{
			"Tony", "staff", 0, 0,
		}
		ub := &UserBuilder{}
		user := ub.Name("Tony").Role("staff").Build()
		assert.Equal(t, user, want)
	})
}
