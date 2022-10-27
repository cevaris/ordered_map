//go:build go1.18
// +build go1.18

package main

import (
	"fmt"

	"github.com/cevaris/ordered_map"
)

func GetAndSetExampleG() {
	// Init new OrderedMap
	om := ordered_map.NewOrderedMapG[string, int]()

	// Set key
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	om.Set("d", 4)

	// Same interface as builtin map
	if val, ok := om.Get("b"); ok {
		// Found key "b"
		fmt.Println(val)
	}

	// Delete a key
	om.Delete("c")

	// Failed Get lookup becase we deleted "c"
	if _, ok := om.Get("c"); !ok {
		// Did not find key "c"
		fmt.Println("c not found")
	}
}

func IteratorExampleG() {
	n := 100
	om := ordered_map.NewOrderedMapG[int, string]()

	for i := 0; i < n; i++ {
		// Insert data into OrderedMap
		om.Set(i, fmt.Sprintf("%d", i*i))
	}

	// Iterate though values
	// - Values iteration are in insert order
	// - Returned in a key/value pair struct
	iter := om.IterFunc()
	for kv, ok := iter(); ok; kv, ok = iter() {
		fmt.Println(kv, kv.Key, kv.Value)
	}
}

type MyStructG struct {
	a int
	b float64
}

func CustomStructG() {
	om := ordered_map.NewOrderedMapG[string, *MyStructG]()
	om.Set("one", &MyStructG{1, 1.1})
	om.Set("two", &MyStructG{2, 2.2})
	om.Set("three", &MyStructG{3, 3.3})

	fmt.Println(om)
	// Ouput: OrderedMap[one:&{1 1.1},  two:&{2 2.2},  three:&{3 3.3}]

}

// func main() {
// 	GetAndSetExampleG()
// 	IteratorExampleG()
// 	CustomStructG()
// }
