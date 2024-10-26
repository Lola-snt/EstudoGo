package arrayVSslice

import (
	"math/rand"
)

type Point struct {
	X float64
	Y float64
}

func newPoint() Point {
	x := rand.Float64()
	y := rand.Float64()
	return Point{
		X: x,
		Y: y,
	}
}

func AddArray() [1000]Point {

	const size = 1000
	var arr [size]Point

	for j := 0; j < size; j++ {
		arr[j] = newPoint()
	}

	return arr
}
func AddSlice() []Point {

	const size = 1000
	var arr []Point

	for j := 0; j < size; j++ {
		arr = append(arr, newPoint())
	}

	return arr
}
