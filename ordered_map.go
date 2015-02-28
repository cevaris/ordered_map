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

type KeyValuePair struct {
	key string
	value int
}

func (k *KeyValuePair) String() string {
	return fmt.Sprintf("%v:%v", k.key, k.value)
}

type Node struct {
	Prev *Node
	Curr *Node
	Key string
}

func NewRootNode() *Node {
	root := &Node{}
	root.Prev = root
	root.Curr = root
	return root
}

func NewNode(prev *Node, curr *Node, key string) *Node {
	return &Node{Prev: prev, Curr: curr, Key:key}
}

func NewOrderedMap(args []*KeyValuePair) *OrderedMap {
	om := &OrderedMap{
		store: make(map[string]int),
		mapper: make(map[string]*Node),
		root: NewRootNode(),
	}
	om.update(args)
	return om
}

func (om *OrderedMap) update(args []*KeyValuePair) {
	for _, pair := range args {
		om.Set(pair.key, pair.value)
	}
}

func (om *OrderedMap) Set(key string, value int) {
	if _, ok := om.store[key]; ok == false {
		root := om.root
		last := root.Prev
		last.Curr = NewNode(last, root, key)
		root.Prev = last.Curr
		om.mapper[key] = last.Curr
		fmt.Println(root.Prev.Key, root.Curr.Key)
	}
	om.store[key] = value
}

func (om *OrderedMap) Get() {
}

func (n *Node) String() string {
	var buffer bytes.Buffer
	root := n
	curr := root.Curr
	for curr != root {
		buffer.WriteString(fmt.Sprintf("%s, ", curr.Key))
		// curr = curr.Curr
	}
	return buffer.String()
	// if n.Prev == nil || n.Curr == nil {
	// 	return fmt.Sprintf("{Prev:%v,Curr:%v,Key:%s}",
	// 	n.Prev, n.Curr, n.Key)
	// } else {
	// 	return "<EMPTY>"
	// }
}

func (om *OrderedMap) String() string {
	// var buffer bytes.Buffer
	// return fmt.Println(buffer.String())
	return fmt.Sprintf("\nStore:%v\nHead:%v\n%v", om.store, om.root, om.mapper)
}

func (om *OrderedMap) Iter() <-chan string {
	keys := make(chan string)
	go func(){
		defer close(keys)
		root := om.root
		var curr *Node = root.Curr
		for curr == root {
			keys <-curr.Key
			curr = curr.Curr
		}
	}()
	return keys
}

func (om1 *OrderedMap) Compare(om2 *OrderedMap) bool {
	return true
}


