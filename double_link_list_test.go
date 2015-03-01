package ordered_map

import (
	"testing"
)

func testStringData() []string {
	var data []string = make([]string, 5)
	data[0] = "test0"
	data[1] = "test1"
	data[2] = "test2"
	data[3] = "test3"
	data[4] = "test4"
	return data
}


func TestAddData(t *testing.T) {
	expected := testStringData()

	ls := NewRootNode()
	if ls == nil {
		t.Error("Failed to create LinkList")
	}

	for _, v := range expected {
		ls.add(v)
	}

	index := 0
	for v := range ls.Iter() {
		if v != expected[index] {
			t.Error("Failed insert of args:", v, expected[index])
		}
		index++
	}
}
