//go:build go1.18
// +build go1.18

package ordered_map

import (
	"fmt"
)

type OrderedMapG[K comparable, V any] struct {
	mapper map[K]*nodeG[K, V]
	root   *nodeG[K, V]
}

func NewOrderedMapG[K comparable, V any]() *OrderedMapG[K, V] {
	om := &OrderedMapG[K, V]{
		mapper: make(map[K]*nodeG[K, V]),
		root:   newRootNodeG[K, V](),
	}
	return om
}

func NewOrderedMapGWithArgs[K comparable, V any](args []*KVPairG[K, V]) *OrderedMapG[K, V] {
	om := NewOrderedMapG[K, V]()
	om.update(args)
	return om
}

func (om *OrderedMapG[K, V]) update(args []*KVPairG[K, V]) {
	for _, pair := range args {
		om.Set(pair.Key, pair.Value)
	}
}

func (om *OrderedMapG[K, V]) Set(key K, value V) {
	if n, ok := om.mapper[key]; ok {
		n.value = value
	}

	root := om.root
	last := root.prev
	n := newNodeG(last, root, key, value)
	last.next = n
	root.prev = n
	om.mapper[key] = n
}

func (om *OrderedMapG[K, V]) Get(key K) (value V, ok bool) {
	if n, ok := om.mapper[key]; ok {
		return n.value, true
	}
	return
}

func (om *OrderedMapG[K, V]) Delete(key K) {
	n, ok := om.mapper[key]
	if ok {
		n.prev.next = n.next
		n.next.prev = n.prev
		delete(om.mapper, key)
	}
}

func (om *OrderedMapG[K, V]) String() string {
	builder := make([]string, len(om.mapper))

	var index int = 0
	iter := om.IterFunc()
	for kv, ok := iter(); ok; kv, ok = iter() {
		val, _ := om.Get(kv.Key)
		builder[index] = fmt.Sprintf("%v:%v", kv.Key, val)
		index++
	}
	return fmt.Sprintf("OrderedMap%v", builder)
}

func (om *OrderedMapG[K, V]) Iter() <-chan *KVPairG[K, V] {
	println("Iter() method is deprecated!. Use IterFunc() instead.")
	return om.UnsafeIter()
}

/*
Beware, Iterator leaks goroutines if we do not fully traverse the map.
For most cases, `IterFunc()` should work as an iterator.
*/
func (om *OrderedMapG[K, V]) UnsafeIter() <-chan *KVPairG[K, V] {
	keys := make(chan *KVPairG[K, V])
	go func() {
		defer close(keys)
		var curr *nodeG[K, V]
		root := om.root
		curr = root.next
		for curr != root {
			keys <- curr.kvpair()
			curr = curr.next
		}
	}()
	return keys
}

func (om *OrderedMapG[K, V]) IterFunc() func() (*KVPairG[K, V], bool) {
	var curr *nodeG[K, V]
	root := om.root
	curr = root.next
	return func() (kvpair *KVPairG[K, V], ok bool) {
		for curr != root {
			kvpair = curr.kvpair()
			curr = curr.next
			return kvpair, true
		}
		return
	}
}

func (om *OrderedMapG[K, V]) Len() int {
	return len(om.mapper)
}
