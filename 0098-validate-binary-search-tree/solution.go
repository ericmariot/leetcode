/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isValid(root, nil, nil)
}

func isValid(root, left, right *TreeNode) bool {
    if root == nil {
        return true
    }

    if right != nil && root.Val >= right.Val {
        return false
    }

    if left != nil && root.Val <= left.Val {
        return false
    }

    return isValid(root.Left, left, root) && isValid(root.Right, root, right)
}
