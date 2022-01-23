# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution1:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        answer = ListNode()
        pre = answer

        while list1 is not None and list2 is not None:
            node = ListNode()
            if list1.val <= list2.val:
                node.val = list1.val
                list1 = list1.next
            else:
                node.val = list2.val
                list2 = list2.next
            pre.next = node
            pre = pre.next

        while list1 is not None:
            node = ListNode(list1.val)
            list1 = list1.next
            pre.next = node
            pre = pre.next

        while list2 is not None:
            node = ListNode(list2.val)
            list2 = list2.next
            pre.next = node
            pre = pre.next

        return answer.next

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution2:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        answer = ListNode()
        pre = answer

        while list1 is not None and list2 is not None:
            if list1.val <= list2.val:
                node = list1
                list1 = list1.next
            else:
                node = list2
                list2 = list2.next
            pre.next = node
            pre = pre.next

        if list1 is not None:
            pre.next = list1
            pre = pre.next
            list1 = list1.next

        if list2 is not None:
            pre.next = list2
            pre = pre.next
            list2 = list2.next

        return answer.next
