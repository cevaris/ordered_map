//go:build go1.18
// +build go1.18

package ordered_map

type nodeG[K comparable, V any] struct {
	prev  *nodeG[K, V]
	next  *nodeG[K, V]
	key   K
	value V
}

func newRootNodeG[K comparable, V any]() *nodeG[K, V] {
	root := &nodeG[K, V]{}
	root.prev = root
	root.next = root
	return root
}

func newNodeG[K comparable, V any](prev *nodeG[K, V], next *nodeG[K, V], key K, value V) *nodeG[K, V] {
	return &nodeG[K, V]{prev: prev, next: next, key: key, value: value}
}

func (n *nodeG[K, V]) kvpair() *KVPairG[K, V] {
	return &KVPairG[K, V]{Key: n.key, Value: n.value}
}
