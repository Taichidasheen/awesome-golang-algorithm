package main

import "testing"

func TestIntHeap_Push(t *testing.T) {
	var heap IntHeap = []int{1,2,3,4,5}
	heap.Push(6) //6push成功
	heap.PushV2(7) //7push失败
	t.Logf("heap:%+v", heap)
}
