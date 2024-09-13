func reorderList(head *ListNode)  {
    turtle, hare := head, head.Next

    for hare != nil && hare.Next != nil {
        turtle = turtle.Next
        hare = hare.Next.Next
    }

    var prev, reversedCurr *ListNode = nil, turtle.Next
    for reversedCurr != nil {
        next := reversedCurr.Next
        reversedCurr.Next = prev
        prev = reversedCurr
        reversedCurr = next
    }
    reversed := prev

    turtle.Next = nil
    curr := head
    
    for curr != nil && reversed != nil {
        next := curr.Next
        revNext := reversed.Next
        curr.Next = reversed
        reversed.Next = next
        curr = next
        reversed = revNext
    }
}

