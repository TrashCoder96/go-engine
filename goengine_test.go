package goengine

import (
	"fmt"
	"testing"
)

//TestSum test represents example of unit test for my go physics engine. In this case test verifies result of sum of two vectors
func TestSum(t *testing.T) {
	var firstVector = Vector{X: 3, Y: 5}
	var secondVector = Vector{X: 2, Y: 3}
	sumVector := firstVector.Add(secondVector)
	if sumVector.X != 5 {
		t.Errorf("X was incorrect, got: %f, want: %d.", sumVector.X, 5)
	}
	if sumVector.Y != 8 {
		t.Errorf("Y incorrect, got: %f, want %d.", sumVector.Y, 8)
	}
}

func TestRunEngine(t *testing.T) {
	world := World{}
	world.AddParticle(Vector{X: 1, Y: 1, Z: 1})
	world.particles[0].SetAcceleration(Vector{0.5, 0.33, 5.0})
	for i := 0; i < 100; i++ {
		world.DoStep(0.1)
	}
	fmt.Println("Engine worked out")
}
