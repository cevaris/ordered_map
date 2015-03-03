package ordered_map

import (
	"testing"
)


type MyStruct struct {
	a float64
	b bool
}

func testIntStruct() []*KVPair {
	var data []*KVPair = make([]*KVPair, 5)
	data[0] = &KVPair{0, &MyStruct{0.1, true}}
	data[1] = &KVPair{1, &MyStruct{1.1, true}}
	data[2] = &KVPair{2, &MyStruct{2.1, false}}
	data[3] = &KVPair{3, &MyStruct{3.1, true}}
	data[4] = &KVPair{4, &MyStruct{4.1, false}}
	return data
}

func testStringInt() []*KVPair {
	var data []*KVPair = make([]*KVPair, 5)
	data[0] = &KVPair{"test0", 0}
	data[1] = &KVPair{"test1", 1}
	data[2] = &KVPair{"test2", 2}
	data[3] = &KVPair{"test3", 3}
	data[4] = &KVPair{"test4", 4}
	return data
}


func TestSetData(t *testing.T) {
	expected := testStringInt()
	om := NewOrderedMap()
	if om == nil {
		t.Error("Failed to create OrderedMap")
	}

	for _, kvp := range expected {
		om.Set(kvp.Key, kvp.Value)
	}

	if len(om.store) != len(expected) {
		t.Error("Failed insert of args:", om.store, expected)
	}
}

func TestGetData(t *testing.T) {
	data := testStringInt()
	om := NewOrderedMapWithArgs(data)

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

func TestDeleteData(t *testing.T) {
	data := testStringInt()
	om := NewOrderedMapWithArgs(data)

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
	sample := testStringInt()
	om := NewOrderedMapWithArgs(sample)

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





func testString() []string {
	var data []string = make([]string, 5)
	data[0] = "test0"
	data[1] = "test1"
	data[2] = "test2"
	data[3] = "test3"
	data[4] = "test4"
	return data
}


func TestAddData(t *testing.T) {
	expected := testString()

	ls := newRootNode()
	if ls == nil {
		t.Error("Failed to create LinkList")
	}

	for _, v := range expected {
		ls.add(v)
	}

	index := 0
	for v := range ls.iter() {
		if v != expected[index] {
			t.Error("Failed insert of args:", v, expected[index])
		}
		index++
	}
}
