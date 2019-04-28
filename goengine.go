package goengine

import "math"

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

//ScalarProduct function returns value, that represent scalar product both vectors A and B (it can be given by formula Ax*Bx+Ay*By or |A|*|B|*cos(alfa))
func (v *Vector) ScalarProduct(vector Vector) float64 {
	return vector.X*v.X + vector.Y*v.Y + vector.Z*v.Z
}

//Normalize function returns new Vector instance that represents normalized vector, given by formula N = [Ax / length, Ay / length]; where length = sqrt(Ax^2 + Xy^2)
func (v *Vector) Normalize() Vector {
	vectorLength := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	normalizedX := v.X / vectorLength
	normalizedY := v.Y / vectorLength
	normalizedZ := v.Z / vectorLength
	return Vector{X: normalizedX, Y: normalizedY, Z: normalizedZ}
}

//VectorProduct function returns new Vector instance, that represents result of vector product of vectors, given by formula
//VP(A, B) = [Ay*Bz - Az*By, Az*Bx - Ax*Bz, Ax*By - Ay*Bx]
func (v *Vector) VectorProduct(vector Vector) Vector {
	return Vector{X: v.Y*vector.Z - v.Z*vector.X, Y: v.Z*vector.X - v.X*vector.Z, Z: v.X*vector.Y - v.Y*vector.X}
}
