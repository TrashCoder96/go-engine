package goengine

import (
	"math"
)

/*vector section*/

//Vector struct represents geometric element of vector space
type Vector struct {
	X float64
	Y float64
	Z float64
}

//Add function initializes new Vector instance, that represents sum of origin and added vectors
func (v *Vector) Add(addedVector Vector) Vector {
	return Vector{X: v.X + addedVector.X, Y: v.Y + addedVector.Y, Z: v.Z + addedVector.Z}
}

//Subtract function initializes new Vector instance, that represents subtruction of subtracted vector from origin vector
func (v *Vector) Subtract(subtractedVector Vector) Vector {
	return Vector{X: v.X - subtractedVector.X, Y: v.Y - subtractedVector.Y, Z: v.Z - subtractedVector.Z}
}

//ComponentProduct function initializes new Vector, that represents component product calculated by multiplying component of each vector by component of other vector
func (v *Vector) ComponentProduct(vector Vector) Vector {
	return Vector{X: v.X * vector.X, Y: v.Y * vector.Y, Z: v.Z * vector.Z}
}

//ScalarProduct function returns value, that represent scalar product both vectors A and B (it can be given by formula Ax*Bx+Ay*By+Az*Bz or |A|*|B|*cos(alfa))
func (v *Vector) ScalarProduct(vector Vector) float64 {
	return vector.X*v.X + vector.Y*v.Y + vector.Z*v.Z
}

//Normalize function returns new Vector instance that represents normalized vector, given by formula N = [Ax / length, Ay / length]; where length = sqrt(x^2 + y^2 + z^2)
func (v *Vector) Normalize() Vector {
	vectorLength := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	normalizedX := v.X / vectorLength
	normalizedY := v.Y / vectorLength
	normalizedZ := v.Z / vectorLength
	return Vector{X: normalizedX, Y: normalizedY, Z: normalizedZ}
}

//Scale function returns new Vector that represents scaled vector based on origin vector, given by formula S = [x * number, y * number, z * number]
func (v *Vector) Scale(number float64) Vector {
	return Vector{X: v.X * number, Y: v.Y * number, Z: v.Z * number}
}

//VectorProduct function returns new Vector instance, that represents result of vector product of vectors, given by formula
//VP(A, B) = [Ay*Bz - Az*By, Az*Bx - Ax*Bz, Ax*By - Ay*Bx]
func (v *Vector) VectorProduct(vector Vector) Vector {
	return Vector{X: v.Y*vector.Z - v.Z*vector.X, Y: v.Z*vector.X - v.X*vector.Z, Z: v.X*vector.Y - v.Y*vector.X}
}

/*end vector section*/

//Particle struct represents simplest unit of physic space without direction
type Particle struct {
	Position       Vector
	LinearVelocity Vector //linearVelocity
	Acceleration   Vector
}

//Integrate function calculate next state of particle, including position and linear velocity
func (particle *Particle) Integrate(duration float64) {
	particle.Position = particle.Position.Add(particle.LinearVelocity.Scale(duration))
	particle.LinearVelocity = particle.LinearVelocity.Add(particle.Acceleration.Scale(duration))
}

//SetAcceleration function sets acceleration vector for particle
func (particle *Particle) SetAcceleration(acc Vector) {
	particle.Acceleration = acc
}

//World struct represents entire game world with its particles
type World struct {
	particles []*Particle
}

//AddParticle function adds new particle into world
func (world *World) AddParticle(position Vector) {
	particle := &Particle{}
	particle.Position = position
	particle.Acceleration = Vector{X: 0, Y: 0, Z: 0}
	particle.LinearVelocity = Vector{X: 0, Y: 0, Z: 0}
	world.particles = append(world.particles, particle)
}

//DoStep function does processing of game world and all its objects
func (world *World) DoStep(duration float64) {
	for _, particle := range world.particles {
		particle.Integrate(duration)
	}
}
