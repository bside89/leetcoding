package poweroftwo

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	n := 1
	expected := true
	result := isPowerOfTwo(n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	n := 16
	expected := true
	result := isPowerOfTwo(n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	n := 3
	expected := false
	result := isPowerOfTwo(n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
