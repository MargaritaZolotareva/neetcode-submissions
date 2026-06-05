/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    occs := make(map[*ListNode]bool)
    node := head

    for node != nil {
        if _, ok := occs[node]; ok {
            return true
        }
        occs[node] = true
        node = node.Next
    }

    return false
}
































    // slow := head
    // fast := head

    // for fast != nil && fast.Next != nil {
    //     slow = slow.Next
    //     fast = fast.Next.Next

    //     if (slow == fast) {
    //         return true
    //     }
    // }

    // return false