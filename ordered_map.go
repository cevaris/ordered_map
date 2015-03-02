package ordered_map

import (
	"fmt"
)

// Taken from
// https://github.com/twitter/commons/blob/master/src/python/twitter/common/collections/ordereddict.py
// http://code.activestate.com/recipes/576693/
// https://github.com/nicklockwood/OrderedDictionary/blob/master/OrderedDictionary/OrderedDictionary.m
// http://golang.org/src/sort/sort.go?s=5371:5390#L223
// http://pymotw.com/2/collections/ordereddict.html

type OrderedMap struct {
	store map[interface{}]interface{}
	mapper map[interface{}]*Node
	root *Node
}

type KVPair struct {
	Key interface{}
	Value interface{}
}

func (k *KVPair) String() string {
	return fmt.Sprintf("%v:%v", k.Key, k.Value)
}

func NewOrderedMap() *OrderedMap {
	om := &OrderedMap{
		store: make(map[interface{}]interface{}),
		mapper: make(map[interface{}]*Node),
		root: NewRootNode(),
	}
	return om
}

func NewOrderedMapWithArgs(args []*KVPair) *OrderedMap {
	om := NewOrderedMap()
	om.update(args)
	return om
}

func (om *OrderedMap) update(args []*KVPair) {
	for _, pair := range args {
		om.Set(pair.Key, pair.Value)
	}
}

func (om *OrderedMap) Set(key interface{}, value interface{}) {
	if _, ok := om.store[key]; ok == false {
		root := om.root
		last := root.Prev
		last.Next = NewNode(last, root, key)
		root.Prev = last.Next
		om.mapper[key] = last.Next
		// fmt.Println(key, value, last.Key, last.Next.Key)
	}
	om.store[key] = value
}

func (om *OrderedMap) Get(key interface{}) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *OrderedMap) Delete(key interface{}) {
	_, ok := om.store[key]
	if ok {
		delete(om.store, key)
	}
	root, rootFound := om.mapper[key]
	if rootFound {
		prev := root.Prev
		next := root.Next
		prev.Next = next
		next.Prev = prev
	}
}



func (om *OrderedMap) String() string {
	builder := make ([]string, len(om.store))

	var index int = 0
	for k := range om.Iter() {
		val, _ := om.Get(k.Key)
		builder[index] = fmt.Sprintf("%v:%v, ", k.Key, val)
		index++
	}
	return fmt.Sprintf("OrderedMap%v", builder)
}

func (om *OrderedMap) Iter() <-chan *KVPair {
	keys := make(chan *KVPair)
	go func(){
		defer close(keys)
		var curr *Node
		root := om.root
		curr = root.Next
		for curr != root {
			v, _ := om.store[curr.Value.(string)]
			keys <- &KVPair{curr.Value.(string), v}
			curr = curr.Next
		}
	}()
	return keys
}

func (om1 *OrderedMap) Compare(om2 *OrderedMap) bool {
	return true
}

func (kv1 *KVPair) Compare(kv2 *KVPair) bool {
	return kv1.Key == kv2.Key && kv1.Value == kv2.Value
}


