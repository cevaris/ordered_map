package ordered_map

import (
	"fmt"
	"bytes"
)

type Node struct {
	Prev *Node
	Next *Node
	Value interface{}
}

func NewRootNode() *Node {
	root := &Node{}
	root.Prev = root
	root.Next = root
	return root
}


func NewNode(prev *Node, next *Node, key interface{}) *Node {
	return &Node{Prev: prev, Next: next, Value:key}
}

func (n *Node) add(value string) {
	root := n
	last := root.Prev
	last.Next = NewNode(last, n, value)
	root.Prev = last.Next
}

func (n *Node) String() string {
	var buffer bytes.Buffer
	if n.Value == "" {
		// Need to sentinel
		var curr *Node
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

func (n *Node) Iter() <-chan string {
	keys := make(chan string)
	go func(){
		defer close(keys)
		var curr *Node
		root := n
		curr = root.Next
		for curr != root {
			keys <- curr.Value.(string)
			curr = curr.Next
		}
	}()
	return keys
}
