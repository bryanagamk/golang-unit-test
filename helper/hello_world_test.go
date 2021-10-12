package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	persons := []struct {
		name string
	}{
		{
			name: "Bro",
		},
		{
			name: "Bryan",
		},
		{
			name: "Kottama",
		},
		{
			name: "Candra Gupta",
		},
	}

	for _, person := range persons {
		b.Run(person.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(person.name)
			}
		})
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Bro", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bro")
		}
	})

	b.Run("Kottama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kottama")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bro")
	}
}

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
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Bro", func(t *testing.T) {
		result := HelloWorld("Bro")
		require.Equal(t, "Hello Bro", result)
	})

	t.Run("Agam", func(t *testing.T) {
		result := HelloWorld("Agam")
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

	result := HelloWorld("Bro")
	require.Equal(t, "Hello Bro", result, "Result must be Hello Bro")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bro")
	require.Equal(t, "Hello Bro", result, "Result must be 'Hello Bro'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Bro")
	assert.Equal(t, "Hello Bro", result, "Result must be 'Hello Bro'")
	fmt.Println("TestHelloWorld with Assertion Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bro")
	if result != "Hello Bro" {
		// unit test failed
		panic("Result is not Hello Bro")
	}
}

func TestHelloWorldAgamError(t *testing.T) {
	result := HelloWorld("Agam")
	if result != "Hello Agam" {
		t.Error("Harusnya Hello Agam")
	}
	fmt.Println("Dieksekusi")
}

func TestHelloWorldAgamFatal(t *testing.T) {
	result := HelloWorld("Agam")
	if result != "Hello Agam" {
		t.Fatal("Harusnya Hello Agam")
	}
	fmt.Println("Tidak Dieksekusi")
}
