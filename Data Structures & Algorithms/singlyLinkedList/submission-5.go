type ListNode struct {
    val int
    next *ListNode
}

func NewListNode (val int, nextNode *ListNode) *ListNode {
    return &ListNode {
        val: val,
        next: nextNode,
    }
}

type LinkedList struct {
    head *ListNode
    tail *ListNode
}

func NewLinkedList() *LinkedList {
    head := NewListNode(-1, nil)
    return &LinkedList{
        head: head,
        tail: head,
    }
}

func (ll *LinkedList) Get(index int) int {
    curr := ll.head.next
    i := 0
    
    for curr != nil {
        if i == index {
            return curr.val
        }
        i++
        curr = curr.next
    }
    return -1
}

func (ll *LinkedList) InsertHead(val int) {
    newNode := NewListNode(val, ll.head.next)
    ll.head.next = newNode
    if newNode.next == nil {
        ll.tail = newNode
    }
}

func (ll *LinkedList) InsertTail(val int) {
    ll.tail.next = NewListNode(val, nil)
    ll.tail = ll.tail.next
}

func (ll *LinkedList) Remove(index int) bool {
    i := 0
    curr := ll.head
    for i < index && curr != nil {
        i++
        curr = curr.next
    }

    if curr != nil && curr.next != nil {
        if curr.next == ll.tail {
            ll.tail = curr
        }
        curr.next = curr.next.next
        return true
    }
    return false
}

func (ll *LinkedList) GetValues() []int {
    curr := ll.head.next
    var res []int
    for curr != nil {
        res = append(res, curr.val)
        curr = curr.next
    }

    return res
}
