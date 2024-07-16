/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    // create a map to save ListNode's
    m := make(map[*ListNode]bool)
    curr := head

    // iterate over the linked list
    for curr != nil {
        // for every item check if it exists in the map
        _, prs := m[curr]
        // if exists return true
        if prs {
            return true
        }

        // if does not exist create it
        m[curr] = true
        curr = curr.Next
    }
    // if the list is over return false
    return false
}
