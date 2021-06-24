package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWoldRequire(t *testing.T) {
	result := HelloWold("Bro")
	require.Equal(t, "Hello Bro", result, "Result must be 'Hello Bro'")
	fmt.Println("TestHelloWold with Require Done")
}

func TestHelloWoldAssertion(t *testing.T) {
	result := HelloWold("Bro")
	assert.Equal(t, "Hello Bro", result, "Result must be 'Hello Bro'")
	fmt.Println("TestHelloWold with Assertion Done")
}

func TestHelloWold(t *testing.T) {
	result := HelloWold("Bro")
	if result != "Hello Bro" {
		// unit test failed
		panic("Result is not Hello Bro")
	}
}

func TestHelloWoldAgam(t *testing.T) {
	result := HelloWold("Agam")
	if result != "Hello Agam" {
		// unit test failed
		panic("Result is not Hello Agam")
	}
}
