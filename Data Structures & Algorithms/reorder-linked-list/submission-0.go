/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	secHalf := listHalf(head)
	secHalf = reverseList(secHalf)
	interleafLists(head, secHalf)
}

func listHalf(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secHalf := slow.Next
	slow.Next = nil
	return secHalf
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode

   for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
   } 
   return prev
}

func interleafLists(l1 *ListNode, l2 *ListNode) {

	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l2Next := l2.Next

		l1.Next = l2
		if l1Next != nil {
			l2.Next = l1Next
		}

		l1 = l1Next
		l2 = l2Next
	}
}