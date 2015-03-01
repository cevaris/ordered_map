package ordered_map

import (
	"testing"
)

func testData() []*KVPair {
	var data []*KVPair = make([]*KVPair, 3)
	data[0] = &KVPair{"test0", 0}
	data[1] = &KVPair{"test1", 1}
	data[2] = &KVPair{"test2", 2}
	return data
}


func TestSetData(t *testing.T) {
	expected := testData()
	om := NewOrderedMap(expected)
	if om == nil {
		t.Error("Failed to create OrderedMap")
	}
	if len(om.store) != len(expected) {
		t.Error("Failed insert of args:", om.store, expected)
	}
}


func TestIterator(t *testing.T) {
	sample := testData()
	om := NewOrderedMap(sample)

	iter := om.Iter()
	if iter == nil {
		t.Error("Failed to create OrderedMap")
	}

	var index int = 0
	for k := range iter {
		expected := sample[index]
		if !k.Compare(expected) {
			t.Error(expected, k)
		}
		index++
	}
}
