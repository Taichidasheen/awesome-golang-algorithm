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
	if node == nil {
		return node
	}
	if node.Next == nil {
		return node
	}
	//last指向新的反转链表的头节点
	last := reverseListRecursion(node.Next)
	node.Next.Next = node
	node.Next = nil
	return last
}

var successor *ListNode

//reverseListRecursionTopN 递归方式反转链表前N个节点
func reverseListRecursionTopN(node *ListNode, topN int64) *ListNode {
	if topN == 1 {
		successor = node.Next
		return node
	}
	//last指向新的反转链表的头节点
	last := reverseListRecursionTopN(node.Next, topN-1)
	node.Next.Next = node
	node.Next = successor
	return last
}

//reverseListRecursionTopN 递归方式反转链表a-b区间的节点
func reverseListRecursionBetween(node *ListNode, a int64, b int64) *ListNode {

	if a == 1 {
		return reverseListRecursionTopN(node, b)
	}
	halfTail := node
	i := a
	for i > 2 {
		halfTail = node.Next
		i = i - 1
	}
	//last指向新的反转链表的头节点
	last := reverseListRecursionTopN(halfTail.Next, b-a+1)
	halfTail.Next = last
	return node
}
