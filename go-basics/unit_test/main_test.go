package main_test

import (
	"testing"
	"unit_test"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    int
		y    int
		want int
	}{
		{
			name: "1+2=3", 
			x: 1, 
			y: 2, 
			want: 3,
		},
		{
			name: "2+3=5",
			x: 2,
			y: 3,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.Add(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    int
		y    int
		want float32
	}{
		{
			name: "1/2=0.5",
			x: 1,
			y: 2,
			want: 0.5,
		},
		{
			name: "3/0=0",
			x: 3,
			y: 0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.Divide(tt.x, tt.y)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
