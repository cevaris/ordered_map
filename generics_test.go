//go:build go1.18
// +build go1.18

package ordered_map

import (
	"testing"
)

func BenchmarkSet_OrderedMap(b *testing.B) {
	om := NewOrderedMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
}

func BenchmarkSet_OrderedMapG(b *testing.B) {
	om := NewOrderedMapG[int, int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
}

func BenchmarkGet_OrderedMap(b *testing.B) {
	om := NewOrderedMap()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		om.Get(i)
	}
}

func BenchmarkGet_OrderedMapG(b *testing.B) {
	om := NewOrderedMapG[int, int]()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		om.Get(i)
	}
}

func BenchmarkDelete_OrderedMap(b *testing.B) {
	om := NewOrderedMap()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		om.Delete(i)
	}
}

func BenchmarkDelete_OrderedMapG(b *testing.B) {
	om := NewOrderedMapG[int, int]()
	for i := 0; i < b.N; i++ {
		om.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		om.Delete(i)
	}
}
