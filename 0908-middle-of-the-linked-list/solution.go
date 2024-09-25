/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    t, h := head, head
    
    for h != nil && h.Next != nil {
        t = t.Next
        h = h.Next.Next
    }
    
    return t
}
