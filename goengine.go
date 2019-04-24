package goengine

//Vector struct represents geometric element of vector space
type Vector struct {
	X int64
	Y int64
}

//Add function initializes new Vector instance, that represents sum of origin and added vectors
func (v *Vector) Add(addedVector Vector) Vector {
	return Vector{X: v.X + addedVector.X, Y: v.Y + addedVector.Y}
}

//Subtract function initializes new Vector instance, that represents subtruction of subtracted vector from origin vector
func (v *Vector) Subtract(subtractedVector Vector) Vector {
	return Vector{X: v.X - subtractedVector.X, Y: v.Y - subtractedVector.Y}
}
