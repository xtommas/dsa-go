package main

import (
	"testing"
)

func expect(name string, result, want interface{}, t *testing.T) {
	if result != want {
		t.Errorf("test %s got %d, want %d", name, result, want)
	}
}
