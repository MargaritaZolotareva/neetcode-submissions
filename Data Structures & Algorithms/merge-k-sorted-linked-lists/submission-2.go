/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if (len(lists) == 0) {
        return nil
    }
    resList := lists[0]

    for i := 1; i < len(lists); i++ {
        resList = mergeLists(resList, lists[i])
    }

    return resList
}

func mergeLists(list1 *ListNode, list2 *ListNode) *ListNode {
    res := &ListNode{Val: 0}
    tail := res

    for list1 != nil && list2 != nil {
        if (list1.Val < list2.Val) {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }

        tail = tail.Next
    }

    if (list1 != nil) {
        tail.Next = list1
    }

    if (list2 != nil) {
        tail.Next = list2
    }

    return res.Next
}



















    // iteration
    // res := &ListNode{Val: 0}
    // cur := res

    // for {
    //     minNode := -1

    //     for i := range lists {
    //         if lists[i] == nil {
    //             continue
    //         }
    //         if (minNode == -1 || lists[i].Val < lists[minNode].Val) {
    //             minNode = i
    //         }
    //     }

    //     if (minNode == -1) {
    //         break
    //     }
        
    //     cur.Next = lists[minNode]
    //     lists[minNode] = lists[minNode].Next
    //     cur = cur.Next
    // }

    // return res.Next