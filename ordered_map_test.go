package ordered_map

import (
	"testing"
)

func testData() []*KVPair {
	var data []*KVPair = make([]*KVPair, 5)
	data[0] = &KVPair{"test0", 0}
	data[1] = &KVPair{"test1", 1}
	data[2] = &KVPair{"test2", 2}
	data[3] = &KVPair{"test3", 3}
	data[4] = &KVPair{"test4", 4}
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

func TestGetData(t *testing.T) {
	data := testData()
	om := NewOrderedMap(data)

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

	t.Log(om)

}

func TestDeleteData(t *testing.T) {
	data := testData()
	om := NewOrderedMap(data)

	testKey := data[2].Key

	// First check to see if exists
	_, ok := om.Get(testKey)
	if !ok {
		t.Error("Key/Value not found in OrderedMap")
	}

	// Delete key
	om.Delete(testKey)

	// Test to see if removed
	_, ok2 := om.Get(testKey)
	if ok2 {
		t.Error("Key/Value was not deleted")
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
