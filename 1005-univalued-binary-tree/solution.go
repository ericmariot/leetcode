/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    return dfs(root, root.Val)
}

func dfs(root *TreeNode, init int) bool {
    if root == nil {
        return true
    }

    if root.Val != init {
        return false
    }

    return dfs(root.Left, init) && dfs(root.Right, init)
}
