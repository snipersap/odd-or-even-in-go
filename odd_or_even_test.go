package main

import "testing"

func TestNewRange(t *testing.T) {
	numberRange := newRange()
	length := len(numberRange)
	if length == 0 {
		t.Errorf("Expected length of number range > 0, found:%v", length)
	}
	if length != 11 {
		t.Errorf("Expected length of number range = 10, found:%v", length)
	}
}
