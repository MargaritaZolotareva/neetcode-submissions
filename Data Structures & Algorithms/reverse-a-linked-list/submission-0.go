/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var prevNode *ListNode
    curr := head
    for curr != nil {
        nextNode := curr.Next
        curr.Next = prevNode
        prevNode = curr
        curr = nextNode
    }

    return prevNode
}
