/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultNode ListNode = ListNode{}
	currentNode := &resultNode
	dx := 0
	for {
		if l1 == nil && l2 == nil && dx == 0 {
			break
		}
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
		n := dx
		if l1 != nil {
			n += l1.Val
		}
		if l2 != nil {
			n += l2.Val
		}
		currentNode.Val = n % 10
		dx = n / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return resultNode.Next
}

// @lc code=end
