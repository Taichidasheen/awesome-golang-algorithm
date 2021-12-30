package main


type IntHeap []int

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
	//h = append(h, x.(int))//如果不使用指针形式的方法接收者，push操作将不生效
}

func (h IntHeap) PushV2(x interface{}) {
	h = append(h, x.(int))
}


func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}