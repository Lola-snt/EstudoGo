package arrayVSslice_test

import (
	"leetcode/arrayVSslice"
	"sort"
	"testing"
)

var result []arrayVSslice.Point

func BenchmarkArray(b *testing.B) {
	var r [1000]arrayVSslice.Point
	for i := 0; i < b.N; i++ {
		r = arrayVSslice.AddArray()
	}
	result = r[:]
	sort.Slice(result, func(i, j int) bool {
		return r[i].X < r[j].Y
	})
}

func BenchmarkSlice(b *testing.B) {
	var r []arrayVSslice.Point

	for i := 0; i < b.N; i++ {
		r = arrayVSslice.AddSlice()
	}
	result = r
	sort.Slice(result, func(i, j int) bool {
		return r[i].X < r[j].Y
	})
}
