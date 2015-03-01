package ordered_map

import (
	"fmt"
	"bytes"
)

// Taken from
// https://github.com/twitter/commons/blob/master/src/python/twitter/common/collections/ordereddict.py
// http://code.activestate.com/recipes/576693/
// https://github.com/nicklockwood/OrderedDictionary/blob/master/OrderedDictionary/OrderedDictionary.m
// http://golang.org/src/sort/sort.go?s=5371:5390#L223
type OrderedMap struct {
	store map[string]int
	mapper map[string]*Node
	root *Node
}

type KVPair struct {
	Key string
	Value int
}

func (k *KVPair) String() string {
	return fmt.Sprintf("%v:%v", k.Key, k.Value)
}

type Node struct {
	Prev *Node
	Next *Node
	Key string
}

func NewRootNode() *Node {
	root := &Node{}
	root.Prev = root
	root.Next = root
	return root
}

func NewNode(prev *Node, next *Node, key string) *Node {
	return &Node{Prev: prev, Next: next, Key:key}
}

func NewOrderedMap(args []*KVPair) *OrderedMap {
	om := &OrderedMap{
		store: make(map[string]int),
		mapper: make(map[string]*Node),
		root: NewRootNode(),
	}
	om.update(args)
	return om
}

func (om *OrderedMap) update(args []*KVPair) {
	for _, pair := range args {
		om.Set(pair.Key, pair.Value)
	}
}

func (om *OrderedMap) Set(key string, value int) {
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

func (om *OrderedMap) Get() {
}

func (n *Node) String() string {
	var buffer bytes.Buffer
	if n.Key == "" {
		// Need to be root
		var curr *Node
		root := n
		curr = root.Next
		for curr != root {
			buffer.WriteString(fmt.Sprintf("%s, ", curr.Key))
			curr = curr.Next
		}
	} else {
		// Else, print pointer value
		buffer.WriteString(fmt.Sprintf("%p, ", &n))
	}
	return buffer.String()
}

func (om *OrderedMap) String() string {
	// var buffer bytes.Buffer
	// return fmt.Println(buffer.String())
	return fmt.Sprintf("\nStore:%v\nHead:%v\n%v", om.store, om.root, om.mapper)
}

func (om *OrderedMap) Iter() <-chan *KVPair {
	keys := make(chan *KVPair)
	go func(){
		defer close(keys)
		var curr *Node
		root := om.root
		curr = root.Next
		for curr != root {
			v, _ := om.store[curr.Key]
			keys <- &KVPair{curr.Key, v}
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


