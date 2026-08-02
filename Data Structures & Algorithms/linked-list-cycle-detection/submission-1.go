/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
   for fast, slow := head.Next, head; fast != nil && fast.Next != nil; slow, fast = slow.Next, fast.Next.Next {
		if fast == slow {
			return true
		}
   }

   return false
}
