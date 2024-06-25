package _map

import "testing"

/*BenchmarkMapInitWithoutSize
BenchmarkMapInitWithoutSize-16    	1000000000	         0.1793 ns/op
BenchmarkMapInitWithSize
BenchmarkMapInitWithSize-16       	1000000000	         0.08701 ns/op

*/

func BenchmarkMapInitWithoutSize(b *testing.B) {
	mapping := make(map[int]int)

	for i := 0; i < 1_000_000; i++ {
		mapping[i] = i
	}
}

func BenchmarkMapInitWithSize(b *testing.B) {
	mapping := make(map[int]int, 1_000_000)

	for i := 0; i < 1_000_000; i++ {
		mapping[i] = i
	}
}
