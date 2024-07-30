package main

import (
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name   string
		grades []float64
		want   float64
	}{
		{"Test case 1", []float64{90, 80, 70, 100}, 85},
		{"Test case 2", []float64{50, 60, 70, 80, 90}, 70},
		{"Test case 3", []float64{100, 100, 100}, 100},
		{"Test case 4", []float64{0, 0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateAverage(tt.grades)
			if got != tt.want {
				t.Errorf("calculateAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
