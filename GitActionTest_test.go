package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 1 + 3
	got := Add(1, 3)

	if expected != got {
		t.Errorf("expected: %d got: %d", expected, got)
	}
}
