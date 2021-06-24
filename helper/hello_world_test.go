package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Bro",
			request:  "Bro",
			expected: "Hello Bro",
		},
		{
			name:     "Bryan",
			request:  "Bryan",
			expected: "Hello Bryan",
		},
		{
			name:     "Agam",
			request:  "Agam",
			expected: "Hello Agam",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWold(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Bro", func(t *testing.T) {
		result := HelloWold("Bro")
		require.Equal(t, "Hello Bro", result)
	})

	t.Run("Agam", func(t *testing.T) {
		result := HelloWold("Agam")
		require.Equal(t, "Hello Agam", result)
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum Unit Test")
	// eksekusi test semua testing dalam package
	m.Run()
	// after
	fmt.Println("Setelah Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWold("Bro")
	require.Equal(t, "Hello Bro", result, "Result must be Hello Bro")
}

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
