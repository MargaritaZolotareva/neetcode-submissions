/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    prevLeft, left, right := head, head, head
    for i := 0; i < n; i++ {
        right = right.Next
    }

    for right != nil {
        prevLeft = left
        left = left.Next
        right = right.Next
    }
    if prevLeft == left {
        head = head.Next
    } else {
        prevLeft.Next = left.Next
    }

    return head
}
