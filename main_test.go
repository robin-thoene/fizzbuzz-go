package main

import (
	"reflect"
	"testing"
)

func TestEx1(t *testing.T) {
	res := Exec(3)
	expected := []string{"1", "2", "Fizz"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("got %v, want %v", res, expected)
	}
}

func TestEx2(t *testing.T) {
	res := Exec(5)
	expected := []string{"1", "2", "Fizz", "4", "Buzz"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("got %v, want %v", res, expected)
	}
}

func TestEx3(t *testing.T) {
	res := Exec(15)
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("got %v, want %v", res, expected)
	}
}
