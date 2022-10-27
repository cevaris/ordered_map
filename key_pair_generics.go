//go:build go1.18
// +build go1.18

package ordered_map

import "fmt"

type KVPairG[K comparable, V any] struct {
	Key   K
	Value V
}

func (k *KVPairG[K, V]) String() string {
	return fmt.Sprintf("%v:%v", k.Key, k.Value)
}
