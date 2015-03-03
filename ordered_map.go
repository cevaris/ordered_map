package ordered_map

import (
	"fmt"
	"bytes"
)


type OrderedMap struct {
	store map[interface{}]interface{}
	mapper map[interface{}]*node
	root *node
}

func NewOrderedMap() *OrderedMap {
	om := &OrderedMap{
		store: make(map[interface{}]interface{}),
		mapper: make(map[interface{}]*node),
		root: newRootNode(),
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
		last.Next = newNode(last, root, key)
		root.Prev = last.Next
		om.mapper[key] = last.Next
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
		var curr *node
		root := om.root
		curr = root.Next
		for curr != root {
			v, _ := om.store[curr.Value]
			keys <- &KVPair{curr.Value, v}
			curr = curr.Next
		}
	}()
	return keys
}




type node struct {
	Prev *node
	Next *node
	Value interface{}
}

func newRootNode() *node {
	root := &node{}
	root.Prev = root
	root.Next = root
	return root
}


func newNode(prev *node, next *node, key interface{}) *node {
	return &node{Prev: prev, Next: next, Value:key}
}

func (n *node) add(value string) {
	root := n
	last := root.Prev
	last.Next = newNode(last, n, value)
	root.Prev = last.Next
}

func (n *node) String() string {
	var buffer bytes.Buffer
	if n.Value == "" {
		// Need to sentinel
		var curr *node
		root := n
		curr = root.Next
		for curr != root {
			buffer.WriteString(fmt.Sprintf("%s, ", curr.Value))
			curr = curr.Next
		}
	} else {
		// Else, print pointer value
		buffer.WriteString(fmt.Sprintf("%p, ", &n))
	}
	return fmt.Sprintf("LinkList[%v]",buffer.String())
}

func (n *node) iter() <-chan string {
	keys := make(chan string)
	go func(){
		defer close(keys)
		var curr *node
		root := n
		curr = root.Next
		for curr != root {
			keys <- curr.Value.(string)
			curr = curr.Next
		}
	}()
	return keys
}







type KVPair struct {
	Key interface{}
	Value interface{}
}

func (k *KVPair) String() string {
	return fmt.Sprintf("%v:%v", k.Key, k.Value)
}

func (kv1 *KVPair) Compare(kv2 *KVPair) bool {
	return kv1.Key == kv2.Key && kv1.Value == kv2.Value
}
