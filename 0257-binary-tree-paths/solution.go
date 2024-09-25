/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    return dfs(root, "")
}

func dfs(node *TreeNode, path string) []string{
    if node == nil {
        return []string{}
    }


    if node.Left == nil && node.Right == nil {
        path = path+strconv.Itoa(node.Val)
        return []string{path}
    }

    path = path+strconv.Itoa(node.Val)+"->"
    dfs(node.Left, path)
    dfs(node.Right, path)

    return append(dfs(node.Left, path), dfs(node.Right, path)...)
}
