/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    dummy := &ListNode{Val: 0, Next: nil}
	for (head != nil){
		next := head.Next
		head.Next = dummy.Next
		dummy.Next = head 
		head = next
	}
	return dummy.Next
}
