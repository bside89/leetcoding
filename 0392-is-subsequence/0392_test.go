package issubsequence

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	s := "abc"
	v := "ahbgdc"
	expected := true
	result := isSubsequence(s, v)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	s := "axc"
	v := "ahbgdc"
	expected := false
	result := isSubsequence(s, v)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}
