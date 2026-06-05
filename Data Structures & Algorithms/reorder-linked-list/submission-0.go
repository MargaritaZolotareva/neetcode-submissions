/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    //recursion
    rec(head, head.Next)
}

func rec(root, curr *ListNode) *ListNode {
    if curr == nil {
        return root
    }
    root = rec(root, curr.Next)

    if root == nil {
        return nil
    }
    var tmp *ListNode
    if root == curr || root.Next == curr {
        curr.Next = nil
    } else {
        tmp = root.Next
        root.Next = curr
        curr.Next = tmp
    }

    return tmp
}