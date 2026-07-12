/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{Val:0, Next:nil}
	current := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val  {
			next := list1.Next
			current.Next = list1
			list1.Next = nil
			list1 = next
		} else {
			next := list2.Next
			current.Next = list2
			list2.Next = nil
			list2 = next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	return dummy.Next
}
