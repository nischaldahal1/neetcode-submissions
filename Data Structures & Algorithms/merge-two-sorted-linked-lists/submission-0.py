# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        last_node = dummy
        while(list1 is not None and list2 is not None):
            if list1.val < list2.val:
                next = list1.next
                last_node.next = list1
                list1.next = None
                last_node = list1
                list1 = next
            else:
                next = list2.next
                last_node.next = list2
                list2.next = None
                last_node = list2
                list2 = next
        if list1 is None:
            last_node.next = list2
        if list2 is None:
            last_node.next = list1
        return dummy.next
            
                
            

