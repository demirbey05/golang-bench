package slice

import "testing"

/*
BenchmarkCreateSliceFromNilSlice: This benchmark measures the time it takes to append elements to a nil slice.
A nil slice has no underlying array, so the append operation may need to allocate memory for a new array when the slice grows.

BenchmarkCreateSliceFromEmptySlice: This benchmark measures the time it takes to append elements to an empty slice.
An empty slice has an underlying array, but its length is 0.
The append operation may need to allocate memory for a new array when the slice grows.

BenchmarkCreateSliceFromSliceWithCapacity: This benchmark measures the time it takes to append elements to a slice that was created with a specific capacity.
The slice has an underlying array with a length of 0, but a capacity equal to b.N.
This means that the append operation does not need to allocate memory for a new array until the number of elements exceeds b.N.

BenchmarkCreateSliceFromMakeEmptySlice: This benchmark measures the time it takes to append elements to a slice that was created with make but without a specific capacity.
The slice has an underlying array with a length of 0.
The append operation may need to allocate memory for a new array when the slice grows.

BenchmarkCreateSliceFromSliceWithLength: This benchmark measures the time it takes to assign values to a slice that was created with a specific length.
The slice has an underlying array with a length equal to b.N, so the assignment operation does not need to allocate memory for a new array.

RESULTS:

BenchmarkCreateSliceFromNilSlice-16             	83156889	        14.19 ns/op
BenchmarkCreateSliceFromEmptySlice-16           	100000000	        11.36 ns/op
BenchmarkCreateSliceFromSliceWithCapacity-16    	448725902	         2.974 ns/op
BenchmarkCreateSliceFromMakeEmptySlice-16       	214379680	        12.32 ns/op
BenchmarkCreateSliceFromSliceWithLength-16      	558077894	         2.256 ns/op
*/

func BenchmarkCreateSliceFromNilSlice(b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}
func BenchmarkCreateSliceFromEmptySlice(b *testing.B) {
	s := []int{}
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}

func BenchmarkCreateSliceFromSliceWithCapacity(b *testing.B) {
	s := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}

func BenchmarkCreateSliceFromMakeEmptySlice(b *testing.B) {
	s := make([]int, 0)
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}

func BenchmarkCreateSliceFromSliceWithLength(b *testing.B) {
	s := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		s[i] = i
	}
}
