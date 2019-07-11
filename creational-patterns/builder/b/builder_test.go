package main

import (
	"reflect"

	"testing"

	"gotest.tools/assert"
)

func TestBuilder(t *testing.T) {
	t.Run("Should create empty struct with default value nil", func(t *testing.T) {
		user, err := BuildUser()
		want := &User{}

		if !reflect.DeepEqual(user, want) {
			t.Errorf("user is not same with want")
		}
		assert.Equal(t, err, nil)
	})

	t.Run("Should set user name", func(t *testing.T) {
		want := "Tony"
		user, err := BuildUser(
			WithName("Tony"),
		)

		assert.Equal(t, user.name, want)
		assert.Equal(t, err, nil)
	})

	t.Run("Should set role as manager", func(t *testing.T) {
		want := User{
			"Tony", "manager", 40000, 60000,
		}
		user, err := BuildUser(
			WithName("Tony"),
			WithRole("manager"),
		)

		assert.Equal(t, user.role, want.role)
		assert.Equal(t, user.minSalary, want.minSalary)
		assert.Equal(t, user.maxSalary, want.maxSalary)
		assert.Equal(t, err, nil)
	})

	t.Run("Should set role as sales", func(t *testing.T) {
		want := User{
			"Tony", "sales", 0, 0,
		}
		user, err := BuildUser(
			WithName("Tony"),
			WithRole("sales"),
		)

		assert.Equal(t, user.role, want.role)
		assert.Equal(t, user.minSalary, want.minSalary)
		assert.Equal(t, user.maxSalary, want.maxSalary)
		assert.Equal(t, err, nil)

	})
}
