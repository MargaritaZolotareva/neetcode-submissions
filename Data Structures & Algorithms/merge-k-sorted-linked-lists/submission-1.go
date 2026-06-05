/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    res := &ListNode{Val: 0}
    cur := res

    for {
        minNode := -1

        for i := range lists {
            if lists[i] == nil {
                continue
            }
            if (minNode == -1 || lists[i].Val < lists[minNode].Val) {
                minNode = i
            }
        }

        if (minNode == -1) {
            break
        }
        
        cur.Next = lists[minNode]
        lists[minNode] = lists[minNode].Next
        cur = cur.Next
    }

    return res.Next
}
