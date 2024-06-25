package _struct

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

/*
BenchmarkCompareWithReflection
BenchmarkCompareWithReflection-16    	 3811074	       292.1 ns/op
BenchmarkCompareWithBytes
BenchmarkCompareWithBytes-16         	  948633	      1243 ns/op
BenchmarkCompareWithCustomFunc
BenchmarkCompareWithCustomFunc-16    	296851962	         4.059 ns/op

Best Method is to use a custom function to compare the struct fields.
*/
func BenchmarkCompareWithReflection(b *testing.B) {
	type data struct {
		Name  string
		slice []float64
	}

	d1 := data{
		Name:  "John Doe",
		slice: []float64{1.1, 2.2, 3.3},
	}
	d2 := data{
		Name:  "John Doe",
		slice: []float64{1.1, 2.2, 3.3},
	}

	for i := 0; i < b.N; i++ {
		if !reflect.DeepEqual(d1, d2) {
			b.Fatal("unexpected")
		}
	}
}

func BenchmarkCompareWithBytes(b *testing.B) {
	type data struct {
		Name  string
		slice []float64
	}

	d1 := data{
		Name:  "John Doe",
		slice: []float64{1.1, 2.2, 3.3},
	}
	d2 := data{
		Name:  "John Doe",
		slice: []float64{1.1, 2.2, 3.3},
	}

	for i := 0; i < b.N; i++ {
		if !bytes.Equal([]byte(fmt.Sprintf("%v", d1)), []byte(fmt.Sprintf("%v", d2))) {
			b.Fatal("unexpected")
		}
	}
}

func BenchmarkCompareWithCustomFunc(b *testing.B) {
	type data struct {
		Name  string
		slice []float64
	}

	d1 := data{
		Name:  "John Doe",
		slice: []float64{1.1, 2.2, 3.3},
	}
	d2 := data{
		Name:  "John Doe",
		slice: []float64{1.1, 2.2, 3.3},
	}

	compare := func(d1, d2 data) bool {
		if d1.Name != d2.Name {
			return false
		}
		if len(d1.slice) != len(d2.slice) {
			return false
		}
		for i := range d1.slice {
			if d1.slice[i] != d2.slice[i] {
				return false
			}
		}
		return true
	}

	for i := 0; i < b.N; i++ {
		if !compare(d1, d2) {
			b.Fatal("unexpected")
		}
	}
}
