package list

import "fmt"

type SinglyLinkedList[T any] struct {
	head     *NodeS[T]
	tail     *NodeS[T]
	len      int
	circular bool
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return new(SinglyLinkedList[T]).Init()
}

// sets l.head, ;.tail = nil && l.len = 0
// does not change l.circular
func (l *SinglyLinkedList[T]) Init() *SinglyLinkedList[T] {
	l.head = nil
	l.tail = nil
	l.len = 0
	return l
}

func (l *SinglyLinkedList[T]) MakeCircular() {
	if l.circular {
		return
	}

	l.circular = true
	if l.head == nil {
		return
	}

	l.tail.next = l.head
}

func (l *SinglyLinkedList[T]) MakeLinear() {
	if !l.circular {
		return
	}

	l.circular = false
	if l.head == nil {
		return
	}

	l.tail.next = nil
}

func (l *SinglyLinkedList[T]) Front() *NodeS[T] {
	return l.head
}

func (l *SinglyLinkedList[T]) Back() *NodeS[T] {
	return l.tail
}

func (l *SinglyLinkedList[T]) Size() int {
	return l.len
}

func (l *SinglyLinkedList[T]) AppendValue(value T) *NodeS[T] {
	return l.Append(&NodeS[T]{
		Value: value,
		list:  l,
	})
}

func (l *SinglyLinkedList[T]) Append(node *NodeS[T]) *NodeS[T] {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
	}

	if l.circular {
		node.next = l.head
	}

	l.tail = node
	l.len += 1

	return node
}

func (l *SinglyLinkedList[T]) PrependValue(value T) *NodeS[T] {
	return l.Prepend(&NodeS[T]{
		Value: value,
		list:  l,
	})
}

func (l *SinglyLinkedList[T]) Prepend(node *NodeS[T]) *NodeS[T] {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
	}

	if l.circular {
		l.tail.next = node
	}

	l.head = node
	l.len += 1

	return node
}

func (l *SinglyLinkedList[T]) Remove(node *NodeS[T]) *T {
	if node == l.head {
		if node.next == nil || node.next == l.head {
			l.Init()
			return &node.Value
		}
		l.head = node.next
		if l.circular {
			l.tail.next = l.head
		}
		l.len -= 1
		return &node.Value
	}

	cur := l.head
	for cur.next != nil {
		if cur.next == node {
			cur.next = node.next
			if node == l.tail {
				l.tail = cur
			}
			l.len -= 1
			break
		}
		cur = cur.next
	}

	if l.len == 0 {
		l.Init()
	}

	return &node.Value
}

func (l *SinglyLinkedList[T]) MoveToFront(node *NodeS[T]) {
	if node == l.head || node == nil || l.head == nil {
		return
	}

	prev := l.head
	cur := l.head.next
	for cur != nil {
		if cur == node {
			switch node {
			case l.tail:
				prev.next = node.next
				l.tail = prev
				node.next = l.head
				l.head = node
				if l.circular {
					l.tail.next = l.head
				}
				return
			default:
				prev.next = node.next
				node.next = l.head
				l.head = node
				if l.circular {
					l.tail.next = l.head
				}
				return
			}
		}
		prev = cur
		cur = cur.next
	}
}

func (l *SinglyLinkedList[T]) MoveToBack(node *NodeS[T]) {
	if node == l.tail || node == nil || l.tail == nil {
		return
	}

	var prev *NodeS[T]
	cur := l.head
	for cur != nil {
		if cur == node {
			switch node {
			case l.head:
				l.head = node.next
				node.next = nil
				l.tail.next = node
				l.tail = node
				if l.circular {
					l.tail.next = l.head
				}
				return
			default:
				prev.next = node.next
				cur = node.next
				if cur == l.tail {
					node.next = nil
					l.tail.next = node
					l.tail = node
					if l.circular {
						l.tail.next = l.head
					}
					return
				}
			}
		}
		prev = cur
		cur = cur.next
	}
}

func (l *SinglyLinkedList[T]) InsertBefore(value T, target *NodeS[T]) *NodeS[T] {
	if target == nil {
		return nil
	}

	node := &NodeS[T]{
		Value: value,
		list:  l,
	}

	if l.head == target {
		return l.Prepend(node)
	}

	cur := l.head
	for cur.next != nil {
		if cur.next == target {
			node.next = cur.next
			cur.next = node
			l.len += 1
			break
		}
		cur = cur.next
	}

	return node
}

func (l *SinglyLinkedList[T]) InsertAfter(value T, target *NodeS[T]) *NodeS[T] {
	if target == nil {
		return nil
	}

	node := &NodeS[T]{
		Value: value,
		list:  l,
	}

	if l.tail == target {
		return l.Append(node)
	}

	cur := l.head
	for cur != nil {
		if cur == target {
			node.next = cur.next
			cur.next = node
			l.len += 1
			break
		}
		cur = cur.next
	}

	return node
}

func (l *SinglyLinkedList[T]) PrintList() {
	var head, tail *T
	if l.head != nil {
		head = &l.head.Value
	}
	if l.tail != nil {
		tail = &l.tail.Value
	}
	fmt.Printf("\nLIST: head %+v tail %+v len %v circular %v\n", head, tail, l.len, l.circular)
}

func (l *SinglyLinkedList[T]) PrintForwards() {
	fmt.Print("\n=== PrintForwards ===\n")
	cur := l.head
	for cur != nil {
		cur.PrintNode()
		cur = cur.next
		if cur == l.head {
			break
		}
	}
}
