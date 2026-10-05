package main

import (
	"reflect"
	"testing"
)

func TestSquareAll(t *testing.T) {
	got := SquareAll([]int{1, 2, 3})
	if !reflect.DeepEqual(got, []int{1, 4, 9}) {
		t.Fatalf("got %v", got)
	}
}

func TestDivide(t *testing.T) {
	if _, err := Divide(1, 0); err == nil {
		t.Fatal("expected error")
	}
	if v, _ := Divide(6, 3); v != 2 {
		t.Fatalf("got %v", v)
	}
}

func TestMap(t *testing.T) {
	got := Map([]int{1, 2}, func(i int) int { return i + 1 })
	if !reflect.DeepEqual(got, []int{2, 3}) {
		t.Fatalf("got %v", got)
	}
}
