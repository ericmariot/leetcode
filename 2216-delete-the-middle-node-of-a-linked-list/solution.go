/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    if head.Next == nil {
        return nil
    }

    b, t, h := head, head, head
    
    for h != nil && h.Next != nil {
        b = t
        t = t.Next
        h = h.Next.Next
    }
    
    right := t.Next
    b.Next = right

    return head
}
