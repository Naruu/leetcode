/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
        return &TreeNode {
            nums[0],
            nil,
            nil,
        }
    }
    max_idx := 0
    max_val := nums[max_idx]
    for i, num := range(nums) {
        if max_val <  num {
            max_idx = i
            max_val = num
        }
    }
    return &TreeNode{
        max_val,
        constructMaximumBinaryTree(nums[:max_idx]),
        constructMaximumBinaryTree(nums[max_idx+1:]),
    }

}