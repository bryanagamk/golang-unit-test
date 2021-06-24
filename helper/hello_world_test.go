package helper

import "testing"

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
