/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    var res [][]int
    lvl := 0
    q := []*TreeNode{root}

    for len(q) > 0 {
        l := len(q)
        for i := 0; i < l; i++ {
            node := q[0]
            q = q[1:]

            if len(res) <= lvl {
                res = append(res, []int{node.Val})
            } else {
                res[lvl] = append(res[lvl], node.Val)
            }

            if node.Left != nil {
                q = append(q, node.Left)
            }

            if node.Right != nil {
                q = append(q, node.Right)
            }
        }

        lvl++
    }

    return res
}

