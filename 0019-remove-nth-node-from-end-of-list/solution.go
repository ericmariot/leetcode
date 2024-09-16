/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l, r := head, head

    for n > 0 {
        r = r.Next
        n--
    }

    if r == nil {
        return l.Next
    }

    for r.Next != nil {
        l = l.Next
        r = r.Next
    }

    l.Next = l.Next.Next

    return head
}
