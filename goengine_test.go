package goengine

import "testing"

//TestSum test represents example of unit test for my go physics engine. In this case test verifies result of sum of two vectors
func TestSum(t *testing.T) {
	var firstVector = Vector{X: 3, Y: 5}
	var secondVector = Vector{X: 2, Y: 3}
	sumVector := firstVector.Add(secondVector)
	if sumVector.X != 5 {
		t.Errorf("X was incorrect, got: %d, want: %d.", sumVector.X, 5)
	}
	if sumVector.Y != 8 {
		t.Errorf("Y incorrect, got: %d, want %d.", sumVector.Y, 8)
	}
}
