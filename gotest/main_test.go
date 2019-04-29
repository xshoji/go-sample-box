package main

import "testing"

func TestReplaceStringSuccess(t *testing.T) {
	result1, result2 := ReplaceString("test[0...10]")
	if result1 != "test" {
		t.Fatalf("failed ReplaceString, %v, %v\n", result1, result2)
	}
	if result2 != "[0...10]" {
		t.Fatalf("failed ReplaceString, %v, %v\n", result1, result2)
	}
}

func TestReplaceStringSuccess2(t *testing.T) {
	result1, result2 := ReplaceString("test")
	if result1 != "test" {
		t.Fatalf("failed ReplaceString, %v, %v\n", result1, result2)
	}
	if result2 != "" {
		t.Fatalf("failed ReplaceString, %v, %v\n", result1, result2)
	}
}
