# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        result_vals = []
        left = l1
        right = l2
        carry = 0
        while left is not None or right is not None or carry != 0:
            if left is not None:
                carry += left.val
                left = left.next

            if right is not None:
                carry += right.val
                right = right.next

            result_vals.append(carry % 10)
            carry = carry // 10

        answer_head = ListNode(result_vals[0])
        head = answer_head
        for val in result_vals[1:]:
            new_node = ListNode(val)
            head.next = new_node
            head = new_node
        return answer_head

