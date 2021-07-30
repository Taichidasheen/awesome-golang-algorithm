package list

import (
	"testing"
)

func Test_reverseListIter(t *testing.T) {
	t.Logf("run test...")
	arr := []int{10, 9, 8}
	originList := UnmarshalListBySlice(arr)
	t.Logf("UnmarshalListBySlice originList:")
	PrintList(originList)
	reverseList := reverseListIter(originList)
	t.Logf("reverseListIter reverseList:")
	PrintList(reverseList)
}

func Test_reverseListRecursion(t *testing.T) {
	t.Logf("run test...")
	arr := []int{10, 9, 8, 7, 6}
	originList := UnmarshalListBySlice(arr)
	t.Logf("UnmarshalListBySlice originList:")
	PrintList(originList)
	reverseList := reverseListRecursion(originList)
	t.Logf("reverseListRecursion reverseList:")
	PrintList(reverseList)
}

func Test_reverseListRecursionTopN(t *testing.T) {
	t.Logf("run test...")
	arr := []int{10, 9, 8, 7, 6}
	originList := UnmarshalListBySlice(arr)
	t.Logf("UnmarshalListBySlice originList:")
	PrintList(originList)
	reverseList := reverseListRecursionTopN(originList, 3)
	t.Logf("reverseListRecursionTopN reverseList:")
	PrintList(reverseList)
}

func Test_reverseListRecursionBetween(t *testing.T) {
	t.Logf("run test...")
	arr := []int{10, 9, 8, 7, 6}
	originList := UnmarshalListBySlice(arr)
	t.Logf("UnmarshalListBySlice originList:")
	PrintList(originList)
	reverseList := reverseListRecursionBetween(originList, 2, 4)
	t.Logf("reverseListRecursionBetween reverseList:")
	PrintList(reverseList)
}
