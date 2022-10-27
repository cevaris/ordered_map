//go:build go1.18
// +build go1.18

package ordered_map

import (
	"testing"
)

func testStringIntG() []*KVPairG[string, int] {
	var data []*KVPairG[string, int] = make([]*KVPairG[string, int], 5)
	data[0] = &KVPairG[string, int]{"test0", 0}
	data[1] = &KVPairG[string, int]{"test1", 1}
	data[2] = &KVPairG[string, int]{"test2", 2}
	data[3] = &KVPairG[string, int]{"test3", 3}
	data[4] = &KVPairG[string, int]{"test4", 4}
	return data
}

func TestSetDataG(t *testing.T) {
	expected := testStringIntG()
	om := NewOrderedMapG[string, int]()
	if om == nil {
		t.Error("Failed to create OrderedMap")
	}

	for _, kvp := range expected {
		om.Set(kvp.Key, kvp.Value)
	}

	if om.Len() != len(expected) {
		t.Error("Failed insert of args:", om.mapper, expected)
	}
}

func TestGetDataG(t *testing.T) {
	data := testStringIntG()
	om := NewOrderedMapGWithArgs(data)

	for _, kvp := range data {
		val, ok := om.Get(kvp.Key)
		if ok && kvp.Value != val {
			t.Error(kvp.Value, val)
		}
	}
	_, ok := om.Get("invlalid-key")
	if ok {
		t.Error("Invalid key was found in OrderedMap")
	}
}

func TestDeleteDataG(t *testing.T) {
	data := testStringIntG()
	om := NewOrderedMapGWithArgs(data)

	testKey := data[2].Key

	// First check to see if exists
	_, ok := om.Get(testKey)
	if !ok {
		t.Error("Key/Value not found in OrderedMap")
	}

	// Assert size equal to "test data size"
	if om.Len() != len(data) {
		t.Error("mapper size is incorrect")
	}

	// Delete key
	om.Delete(testKey)

	// Assert size equal to "test data size" - 1
	if om.Len() != (len(data) - 1) {
		t.Error("mapper size is incorrect")
	}

	// Test to see if removed
	_, ok2 := om.Get(testKey)
	if ok2 {
		t.Error("Key/Value was not deleted")
	}
}

func TestIteratorG(t *testing.T) {
	sample := testStringIntG()
	om := NewOrderedMapGWithArgs(sample)
	iter := om.UnsafeIter()
	if iter == nil {
		t.Error("Failed to create OrderedMap")
	}

	var index int = 0
	for k := range iter {
		expected := sample[index]
		if k.Key != expected.Key || k.Value != expected.Value {
			t.Error(expected, k)
		}
		index++
	}
}

func TestIteratorFuncG(t *testing.T) {
	sample := testStringIntG()
	om := NewOrderedMapGWithArgs(sample)

	iter := om.IterFunc()
	if iter == nil {
		t.Error("Failed to create OrderedMap")
	}

	var index int = 0
	for k, ok := iter(); ok; k, ok = iter() {
		expected := sample[index]
		if k.Key != expected.Key || k.Value != expected.Value {
			t.Error(expected, k)
		}
		index++
	}
}

func TestLenNonEmptyG(t *testing.T) {
	data := testStringIntG()
	om := NewOrderedMapGWithArgs(data)

	if om.Len() != len(data) {
		t.Fatal("Unexpected length")
	}
}

func TestLenEmptyG(t *testing.T) {
	om := NewOrderedMapG[string, int]()

	if om.Len() != 0 {
		t.Fatal("Unexpected length")
	}
}
