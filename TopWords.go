package main

import (
	"fmt"
	"sort"
	"sync"
)

// TopNum is the number of most common words
const TopNum = 5

type KV struct {
	Key string
	Val int
}

// TopN struct that stores the TopNum most common words
type TopN struct {
	Lock  sync.RWMutex
	TopKV [TopNum]KV
}

// NewTopN returns a pointer to the TopN struct and
// initializes the TopKV slice values to the value received
func NewTopN(val int) *TopN {
	t := &TopN{
		Lock:  sync.RWMutex{},
		TopKV: [TopNum]KV{},
	}
	for i, _ := range t.TopKV {
		t.TopKV[i] = KV{
			Val: val,
		}
	}
	return t
}

// Len returns the length of the TopKV slice
func (t *TopN) Len() int {
	return len(t.TopKV)
}

// GetKey returns the key in the index received
func (t *TopN) GetKey(i int) string {
	if i > t.Len() {
		return ""
	}
	t.Lock.RLock()
	defer t.Lock.RUnlock()

	return t.TopKV[i].Key
}

// GetVal returns the value in the index received
func (t *TopN) GetVal(i int) int {
	if i > t.Len() {
		return -1
	}
	t.Lock.RLock()
	defer t.Lock.RUnlock()

	return t.TopKV[i].Val
}

// GetMinMax returns the min & max values
func (t *TopN) GetMinMax() (max int, min int) {
	t.Lock.RLock()
	defer t.Lock.RUnlock()

	max = t.GetVal(0)
	if max == 0 {
		max = 1
	}
	min = t.GetVal(t.Len() - 1)
	if min == 0 {
		min = 1
	}
	return max, min
}

// GetKV returns the KV struct
func (t *TopN) GetKV() [TopNum]KV {
	t.Lock.RLock()
	defer t.Lock.RUnlock()

	return t.TopKV
}

// UpdateKV updates the KV struct in the key, value and, index received
func (t *TopN) UpdateKV(key string, val, i int) {
	if i > t.Len() {
		return
	}
	t.Lock.Lock()
	defer t.Lock.Unlock()

	t.TopKV[i] = KV{
		Key: key,
		Val: val,
	}
}

// UpdateVal updates the value in the index received
func (t *TopN) UpdateVal(val, i int) {
	t.Lock.Lock()
	defer t.Lock.Unlock()

	t.TopKV[i].Val = val
}

// Search searches the key in the TopKV slice, if the key exists, returns the
// key index & val, otherwise, returns the length of the TopKV slice & 0 - O(n)
func (t *TopN) Search(key string) (i int, val int) {
	t.Lock.RLock()
	defer t.Lock.RUnlock()

	for i, _ = range t.TopKV {
		if t.TopKV[i].Key == key {
			return i, t.TopKV[i].Val
		}
	}
	return TopNum, 0
}

// Sort sorts the TopKV slice by descending order - O(n*log(n))
func (t *TopN) Sort() {
	t.Lock.Lock()
	defer t.Lock.Unlock()

	sort.Slice(t.TopKV[:], func(i, j int) bool {
		return t.TopKV[i].Val > t.TopKV[j].Val
	})
}

// ToString prints the TopKV slice - O(n)
func (t *TopN) ToString() string {
	defer t.Lock.RUnlock()
	t.Lock.RLock()

	str := ""
	for _, kv := range t.TopKV {
		str = fmt.Sprintf("%s%s %d\n", str, kv.Key, kv.Val)
	}
	return str
}

// Print prints the TopKV slice - O(n)
func (t *TopN) Print() {
	defer t.Lock.RUnlock()
	t.Lock.RLock()

	for i, kv := range t.TopKV {
		fmt.Printf("\t[%d] %s: %d\n", i, kv.Key, kv.Val)
	}
}
