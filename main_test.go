package main

import (
	"bytes"
	"log"
	"testing"
)

func TestXxx(t *testing.T) {
	buf := bytes.Buffer{}
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Println("This is a test")

	result := add(1, 2)

	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}
