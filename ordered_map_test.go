package ordered_map

import (
	"testing"
)

func testData() []*KeyValuePair {
	var data []*KeyValuePair = make([]*KeyValuePair, 3)
	data[0] = &KeyValuePair{"test0", 0}
	data[1] = &KeyValuePair{"test1", 1}
	data[2] = &KeyValuePair{"test2", 2}
	return data
}


// func TestSetData(t *testing.T) {
// 	expected := testData()
// 	om := NewOrderedMap(expected)
// 	if om == nil {
// 		t.Error("Failed to create OrderedMap")
// 	}
// 	if len(om.store) != len(expected) {
// 		t.Error("Failed insert of args:", om.store, expected)
// 	}
// }


func TestIterator(t *testing.T) {
	sample := testData()
	om := NewOrderedMap(sample)
	// iter := om.Iter()
	// if iter == nil {
	// 	t.Error("Failed to create OrderedMap")
	// }
	t.Log(om.root)
	// for k := range iter {
	// 	t.Log(k)
	// }

}
