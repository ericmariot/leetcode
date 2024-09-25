/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    return dfs(root, val)
}

func dfs(node *TreeNode, val int) *TreeNode {
    if node.Val == val {
        return node
    }

    if node.Left != nil && val < node.Val {
        return dfs(node.Left, val)
    }

    if node.Right != nil && val > node.Val {
        return dfs(node.Right, val)
    }

    return nil
}
