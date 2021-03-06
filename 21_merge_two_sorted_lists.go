/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    answer := &ListNode{0, nil}
    pre := answer

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            pre.Next = list1
            list1 = list1.Next
        } else {
            pre.Next = list2
            list2 = list2.Next
        }
        pre = pre.Next
    }

    for list1 != nil {
        pre.Next = list1
        list1 = list1.Next
        pre = pre.Next
    }

    for list2 != nil {
        pre.Next = list2
        list2 = list2.Next
        pre = pre.Next
    }

    return answer.Next
}