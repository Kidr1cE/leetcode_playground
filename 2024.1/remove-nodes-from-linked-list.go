/*
Date: 2023-12-1
ProblemID: 1661
ProblemName: 找出叠涂元素
*/

package leetcode

/*
*
Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeNodes(head.Next)
	if head.Next != nil && head.Val < head.Next.Val {
		return head.Next
	} else {
		return head
	}
}
