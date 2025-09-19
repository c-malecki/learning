package list

import "fmt"

// todo: clean up or possibly just split into different lists to remove conditional logic

type LinkedList[T any] struct {
	head     *Node[T]
	tail     *Node[T]
	len      int
	double   bool
	circular bool
}

// initialized as a singly linked list
func New[T any]() *LinkedList[T] {
	list := &LinkedList[T]{}
	list.Reset()
	return list
}

// resets list to singly linked list
func (l *LinkedList[T]) Reset() *LinkedList[T] {
	l.head = nil
	l.tail = nil
	l.double = false
	l.circular = false
	l.len = 0
	return l
}

// makes a linked list doubly linked
func (l *LinkedList[T]) MakeDoubly() *LinkedList[T] {
	if l.double {
		return l
	}

	if l.head == nil {
		l.double = true
		return l
	}

	prev := l.head
	cur := l.head
	for cur.next != nil {
		if cur.next == l.head {
			break
		}
		cur = cur.next
		cur.prev = prev
		prev = cur
	}

	l.tail = cur
	if l.circular {
		l.head.prev = l.tail
		l.tail.next = l.head
	}

	l.double = true

	return l
}

// makes a linked list singly linked
func (l *LinkedList[T]) MakeSingly() *LinkedList[T] {
	if !l.double {
		return l
	}

	if l.head == nil {
		l.double = false
		return l
	}

	cur := l.head
	for cur.next != nil {
		cur.prev = nil
		if cur.next == l.head {
			break
		}
		cur = cur.next
	}
	l.tail.prev = nil
	l.double = false

	return l
}

// makes a singly or doubly linked list circular
func (l *LinkedList[T]) MakeCircular() *LinkedList[T] {
	if l.circular {
		return l
	}

	if l.head == nil {
		l.circular = true
		return l
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = l.head

	if l.double {
		l.head.prev = cur
	}
	l.circular = true

	return l
}

// makes a singly or doubly linked list linear
func (l *LinkedList[T]) MakeLinear() *LinkedList[T] {
	if !l.circular {
		return l
	}

	if l.double {
		l.head.prev = nil
		l.tail.next = nil
		return l
	}

	cur := l.head
	for cur.next != nil {
		if cur.next == l.head {
			cur.next = nil
			break
		}
		cur = cur.next
	}

	return l
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.head
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Size() int {
	return l.len
}

func (l *LinkedList[T]) AppendValue(value T) *Node[T] {
	return l.Append(&Node[T]{
		Value: value,
		list:  l,
	})
}

func (l *LinkedList[T]) Append(node *Node[T]) *Node[T] {
	if l.head == nil {
		l.head = node
		l.tail = node
		if l.circular {
			l.head.next = l.head
			if l.double {
				l.head.prev = l.tail
			}
		}
		l.len += 1
		return node
	}

	move := l.tail
	if l.double {
		node.prev = move
	}
	move.next = node
	if l.circular {
		node.next = l.head
	}
	l.tail = node

	if l.double && l.circular {
		l.head.prev = l.tail
	}

	l.len += 1

	return node
}

func (l *LinkedList[T]) PrependValue(value T) *Node[T] {
	return l.Prepend(&Node[T]{
		Value: value,
		list:  l,
	})
}

func (l *LinkedList[T]) Prepend(node *Node[T]) *Node[T] {
	if l.head == nil {
		l.head = node
		l.tail = node
		if l.circular {
			l.head.next = l.head
			if l.double {
				l.head.prev = l.tail
			}
		}
		l.len += 1
		return node
	}

	move := l.head
	l.head = node
	l.head.next = move
	if l.circular {
		l.tail.next = l.head
	}

	if l.double {
		move.prev = l.head
		if l.circular {
			l.head.prev = l.tail
		} else {
			l.head.prev = nil
		}
	}

	l.len += 1

	return node
}

func (l *LinkedList[T]) Remove(node *Node[T]) T {
	if l.head == node {
		if l.len == 1 {
			l.head = nil
			l.tail = nil
			l.len = 0
			return node.Value
		}

		l.head = node.next
		if l.circular {
			l.tail.next = l.head
		}

		if l.double && l.circular {
			l.head.prev = l.tail
		}

		l.len -= 1
		return node.Value
	}

	cur := l.head
	for cur.next != nil {
		if cur.next == node {
			cur.next = node.next
			if l.circular && l.double {
				cur.next.prev = cur
			}
			if node == l.tail {
				l.tail = cur
			}
			l.len -= 1
			break
		}
		cur = cur.next
	}

	return node.Value
}

func (l *LinkedList[T]) MoveToFront(node *Node[T]) {
	if node == nil || l.head == nil || node == l.head {
		return
	}

	move := l.head
	cur := l.head
	for cur.next != nil {
		if cur.next == node && cur == l.head {
			move.next = node.next
			node.next = move
			if l.double {
				move.prev = node
			}
			break
		}

		cur = cur.next
		if cur.next == node {
			cur.next = node.next
			node.next = move
			if l.double {
				move.prev = node
			}
			if node == l.tail {
				l.tail = cur
			}
			break
		}
	}

	l.head = node

	if l.circular {
		l.tail.next = l.head
		if l.double {
			l.head.prev = l.tail
		}
	} else {
		l.tail.next = nil
		l.head.prev = nil
	}
}

func (l *LinkedList[T]) MoveToBack(node *Node[T]) {
	if node == nil || l.tail == nil || node == l.tail {
		return
	}

	if node == l.head {
		l.head = node.next
		cur := l.head
		for cur.next != nil {
			cur = cur.next
			if cur.next == nil || cur.next == node {
				if l.double {
					node.prev = cur
				}
				if l.circular {
					node.next = l.head
				}
				cur.next = node
				break
			}
			cur = cur.next
		}
	} else {
		cur := l.head
		for cur.next != nil {
			if cur.next == node {
				cur.next = node.next
				if l.double {
					cur.next.prev = cur
				}
			}
			if cur.next == l.head {
				break
			}
			cur = cur.next
		}
		cur.next = node
	}

	l.tail = node

	if l.circular {
		l.tail.next = l.head
		if l.double {
			l.head.prev = l.tail
		}
	} else {
		l.tail.next = nil
		l.head.prev = nil
	}
}

func (l *LinkedList[T]) InsertBefore(value T, target *Node[T]) *Node[T] {
	if target == nil {
		return nil
	}

	node := &Node[T]{
		Value: value,
		list:  l,
	}

	if l.head == target {
		return l.Prepend(node)
	}

	cur := l.head
	for cur.next != nil {
		if cur.next == target {
			cur.next = node
			node.next = target
			if l.double {
				target.prev = node
				node.prev = cur
			}
			l.len += 1
			return node
		}
		cur = cur.next
	}

	return nil
}

func (l *LinkedList[T]) InsertAfter(value T, target *Node[T]) *Node[T] {
	if target == nil {
		return nil
	}

	node := &Node[T]{
		Value: value,
		list:  l,
	}

	if l.tail == target {
		return l.Append(node)
	}

	if l.head == target {
		move := l.head.next
		l.head.next = node
		node.next = move
		if l.double && move != nil {
			move.prev = node
		}
		l.len += 1
		return node
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
		if cur == target {
			move := cur.next
			node.next = move
			cur.next = node
			if l.double {
				cur.next.prev = cur
			}
			l.len += 1
			return node
		}
	}

	return nil
}

func (l *LinkedList[T]) PrintList() {
	var head, tail *T
	if l.head != nil {
		head = &l.head.Value
	}
	if l.tail != nil {
		tail = &l.tail.Value
	}
	fmt.Printf("LIST: head %+v tail %+v len %v double %v circular %v\n\n", head, tail, l.len, l.double, l.circular)
}

func (l *LinkedList[T]) PrintForwards() {
	fmt.Print("\n\n=== PrintForwards ===\n")
	cur := l.head
	for cur != nil {
		cur.PrintNode()
		cur = cur.next
		if cur == l.head {
			break
		}
	}
}

func (l *LinkedList[T]) PrintBackwards() {
	fmt.Print("\n\n=== PrintBackwards ===\n")
	if !l.double {
		fmt.Print("\n\nlist is not doubly linked\n")
		return
	}
	cur := l.tail
	for cur != nil {
		cur.PrintNode()
		cur = cur.prev
		if cur == l.tail {
			break
		}
	}
}
