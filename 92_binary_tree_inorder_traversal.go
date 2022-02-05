/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    var answer []int
    var traverse func (node *TreeNode)
    traverse = func (node *TreeNode) {
        if node == nil {
            return
        }
        traverse(node.Left)
        answer = append(answer, node.Val)
        traverse(node.Right)
    }
    traverse(root)
    return answer
}