func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    // Step 1: Find the middle of the list using slow/fast pointers.
    // Starting fast at head.Next ensures slow stops exactly at the end of the first half.
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // Step 2: Split the list into two halves and reverse the second half.
    head2 := slow.Next
    slow.Next = nil // Sever the connection between 1st and 2nd half
    
    head2 = reverseList(head2)

    // Step 3: Merge the two halves in-place without dummy nodes.
    curr1, curr2 := head, head2
    for curr2 != nil {
        next1, next2 := curr1.Next, curr2.Next

        curr1.Next = curr2  // Connect 1st half node to 2nd half node
        curr2.Next = next1  // Connect 2nd half node to next 1st half node

        curr1 = next1
        curr2 = next2
    }
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}