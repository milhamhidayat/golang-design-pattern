package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance_ReturnSingleton(t *testing.T) {
	t.Parallel()
	singleton := GetInstance()
	assert.NotNil(t, singleton)
}

// test multiple call to create instance
// if multiple call, then the second instance must be same with the first instance
func TestGetInstance_MultipleCallsToGetInstance_ReturnsSingleton(t *testing.T) {
	t.Parallel()
	singleton := GetInstance()
	newSingleton := GetInstance()
	assert.Equal(t, singleton, newSingleton)
}

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	assert.Equal(t, instance1, instance2)
}
