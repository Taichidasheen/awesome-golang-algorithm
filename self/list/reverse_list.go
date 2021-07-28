package list

//reverseListIter 迭代方式反转链表
func reverseListIter(node *ListNode) *ListNode {

	var cur, next *ListNode
	cur = node
	next = cur.Next
	cur.Next = nil
	for next != nil {
		nextnext := next.Next
		next.Next = cur
		cur = next
		next = nextnext
	}
	return cur
}

//reverseListRecursion 递归方式反转链表
func reverseListRecursion(node *ListNode) *ListNode {

}
