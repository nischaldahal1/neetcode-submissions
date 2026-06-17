# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0,None)
        current = head
        while(current is not None):
            next_node = current.next
            current.next = dummy.next
            dummy.next = current
            current= next_node
        return dummy.next
            