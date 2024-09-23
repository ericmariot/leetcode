/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    return slices.Equal(leaves(root1), leaves(root2))
}

func leaves(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    }

    return append(leaves(root.Left), leaves(root.Right)...)
}
