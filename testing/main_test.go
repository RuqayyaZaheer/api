package main

import "testing"

var tests = []struct {
	name     string
	divided  float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid", 100.0, 10.0, 10.0, false},
	{"notvalid", 100.0, 0.0, 0.0, true},
	{"nValid", 34.0, 5.0, 6.80, false},
}

func TestDisions(t *testing.T) {

	for _, tt := range tests {
		got, err := divied(tt.divided, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expacted an error but not fond", err)
			}
		} else {
			if err != nil {
				t.Error("did not expected error but found one")
			}

		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}

	}

}

func TestDivide(t *testing.T) {
	_, err := divied(10, 10)

	if err != nil {
		t.Error("Got an error when we should not have")
	}

}

func TestBadDivide(t *testing.T) {
	_, err := divied(10, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")

	}
}
