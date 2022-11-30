package gtypes

import "math"

type Vector2 struct {
	X          float64
	Y          float64
	cachedMag  float64
	cachedMagX float64
	cachedMagY float64
}

func Vector2_One() Vector2 {
	return Vector2{X: .5, Y: .5}
}

func Vector2_Zero() Vector2 {
	return Vector2{X: 0, Y: 0}
}

func Vector2_DistanceBetween(p1 Vector2, p2 Vector2) float64 {
	res := math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
	return res
}

func (v2 *Vector2) Magnitude() float64 {
	if v2.X != v2.cachedMagX || v2.Y != v2.cachedMagY {
		mag := math.Abs(math.Sqrt(math.Pow(v2.X, 2) + math.Pow(v2.Y, 2)))
		v2.cachedMag = mag
		v2.cachedMagX = v2.X
		v2.cachedMagY = v2.Y
		return mag
	} else {
		return v2.cachedMag
	}
}

func (v2 *Vector2) Normalize() *Vector2 {
	sum := math.Abs(v2.X) + math.Abs(v2.Y)
	res := Vector2{X: v2.X / sum, Y: v2.Y / sum}
	return &res
}

func (v2 *Vector2) ChangeMagnitudeBy(factor float64) *Vector2 {
	res := Vector2{X: v2.X * factor, Y: v2.Y * factor}
	return &res
}

func (v2 *Vector2) Add(inVector Vector2) *Vector2 {
	newX := v2.X + inVector.X
	newY := v2.Y + inVector.Y
	return &Vector2{X: newX, Y: newY}
}
