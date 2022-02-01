/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    answer := head
    pre := head

    for pre != nil && pre.Next != nil {
        if pre.Next.Val == pre.Val {
            pre.Next = pre.Next.Next
        } else {
            pre = pre.Next
        }
    }
    return answer
}