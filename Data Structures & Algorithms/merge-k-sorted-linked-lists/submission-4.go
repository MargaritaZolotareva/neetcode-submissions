/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    for len(lists) > 1 {
        var mergedLists []*ListNode
        for i := 0; i < len(lists); i += 2 {
            l1 := lists[i]
            var l2 *ListNode
            if i + 1 < len(lists) {
                l2 = lists[i + 1]
            }
            mergedLists = append(mergedLists, mergeLists(l1, l2))
        }
        lists = mergedLists
    }

    return lists[0]
}

func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{Val: 0}
    curr := res
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }

        curr = curr.Next
    }

    if (l1 != nil) {
        curr.Next = l1
    }
    if (l2 != nil) {
        curr.Next = l2
    }

    return res.Next
}


































// divide & conquer (recursion)
//     return divide(lists, 0, len(lists) - 1)

// }

// func divide(lists []*ListNode, left, right int) *ListNode {
//     if left > right {
//         return nil
//     }
//     if left == right {
//         return lists[left]
//     }

//     mid := (right + left) / 2

//     l1 := divide(lists, left, mid)
//     l2 := divide(lists, mid + 1, right)

//     return conquer(l1, l2)
// }

// func conquer(l1, l2 *ListNode) *ListNode {
//     res := &ListNode{Val: 0}
//     curr := res

//     for l1 != nil && l2 != nil {
//         if l1.Val < l2.Val {
//             curr.Next = l1
//             l1 = l1.Next
//         } else {
//             curr.Next = l2
//             l2 = l2.Next
//         }

//         curr = curr.Next
//     }

//     if (l1 != nil) {
//         curr.Next = l1
//     }
//     if (l2 != nil) {
//         curr.Next = l2
//     }
//     return res.Next
// }

    // merge lists one by one
//     if (len(lists) == 0) {
//         return nil
//     }
//     resList := lists[0]

//     for i := 1; i < len(lists); i++ {
//         resList = mergeLists(resList, lists[i])
//     }

//     return resList
// }

// func mergeLists(list1 *ListNode, list2 *ListNode) *ListNode {
//     res := &ListNode{Val: 0}
//     tail := res

//     for list1 != nil && list2 != nil {
//         if (list1.Val < list2.Val) {
//             tail.Next = list1
//             list1 = list1.Next
//         } else {
//             tail.Next = list2
//             list2 = list2.Next
//         }

//         tail = tail.Next
//     }

//     if (list1 != nil) {
//         tail.Next = list1
//     }

//     if (list2 != nil) {
//         tail.Next = list2
//     }

//     return res.Next
// }
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