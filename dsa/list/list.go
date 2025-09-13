package list

import "fmt"

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

func (l *LinkedList[T]) Print() {
	var head, tail *T
	if l.head != nil {
		head = &l.head.Value
	}
	if l.tail != nil {
		tail = &l.tail.Value
	}
	fmt.Printf("LIST: head %+v tail %+v len %v double %v circular %v\n\n", head, tail, l.len, l.double, l.circular)
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

// makes a singly linked list doubly linked and accounts for circular
func (l *LinkedList[T]) MakeDoubly() *LinkedList[T] {
	if l.double {
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

// makes a doubly linked list singly
func (l *LinkedList[T]) MakeSingly() *LinkedList[T] {
	if !l.double {
		return l
	}

	prev := l.head
	cur := l.head
	for cur.prev != nil {
		if cur.prev == l.head {
			cur.prev = nil
			break
		}
		cur = cur.prev
		prev.prev = nil
		prev = cur
	}

	l.tail = nil
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

func (l *LinkedList[T]) AppendValue(value T) *Node[T] {
	return l.Append(&Node[T]{
		Value: value,
		list:  l,
	})
}

func (l *LinkedList[T]) Append(node *Node[T]) *Node[T] {
	if l.circular {
		return l.insertBackCircular(node)
	}
	return l.insertBackLinear(node)
}

func (l *LinkedList[T]) insertBackLinear(node *Node[T]) *Node[T] {
	if l.head == nil {
		l.head = node
		l.len += 1
		return node
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node

	if l.double {
		node.prev = cur
		l.tail = node
	}

	l.len += 1

	return node
}

func (l *LinkedList[T]) insertBackCircular(node *Node[T]) *Node[T] {
	if l.head == nil {
		l.head = node
		l.head.next = l.head
		if l.double {
			l.head.prev = l.head
		}
		l.len += 1
		return node
	}

	cur := l.head
	for cur.next != nil {
		if cur.next == l.head {
			break
		}
		cur = cur.next
	}
	cur.next = node

	if l.double {
		node.prev = cur
		l.tail = node
		l.head.prev = l.tail
		l.tail.next = l.head
	}

	if !l.double {
		node.next = l.head
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
	if l.circular {
		return l.insertFrontCircular(node)
	}
	return l.insertFrontLinear(node)
}

func (l *LinkedList[T]) insertFrontCircular(node *Node[T]) *Node[T] {
	if l.head == nil {
		l.head = node
		l.head.next = l.head
		l.head.prev = l.head
		l.len += 1
		return node
	}

	move := l.head
	node.next = move
	l.head = node

	if l.double {
		move.prev = l.head
		l.head.prev = l.tail
		l.tail.next = l.head
	}

	l.len += 1

	return node
}

func (l *LinkedList[T]) insertFrontLinear(node *Node[T]) *Node[T] {
	if l.head == nil {
		l.head = node
		l.len += 1
		return node
	}

	move := l.head
	l.head = node
	node.next = move

	if l.double {
		move.prev = l.head
	}

	l.len += 1

	return node
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if l.double {
		l.removeDoublyNode(node)
	} else {
		l.removeSinglyNode(node)
	}
}

func (l *LinkedList[T]) removeSinglyNode(node *Node[T]) {
	if l.head == node {
		if l.len == 1 {
			l.head = nil
			l.len = 0
			return
		}

		l.head = node.next
		if l.circular {
			cur := l.head
			for cur.next != nil {
				if cur.next == node {
					cur.next = l.head
					break
				}
				cur = cur.next
			}
		}
		l.len -= 1
		return
	}

	cur := l.head
	for cur.next != nil {
		if cur.next == node {
			cur.next = node.next
			l.len -= 1
			break
		}
		cur = cur.next
	}
}

func (l *LinkedList[T]) removeDoublyNode(node *Node[T]) {
	if l.head == node {
		if l.len == 1 {
			l.head = nil
			l.tail = nil
			l.len = 0
			return
		}

		l.head = node.next
		if l.circular {
			l.head.prev = l.tail
			l.tail.next = l.head
		} else {
			l.head.prev = nil
		}
		l.len -= 1
		return
	}

	if l.tail == node {
		l.tail = node.prev
		if l.circular {
			l.head.prev = l.tail
			l.tail.next = l.head
		} else {
			l.tail.next = nil
		}
		l.len -= 1
		return
	}

	cur := l.head
	for cur.next != nil {
		if cur.next == node {
			cur.next = node.next
			cur.next.prev = cur
			l.len -= 1
			return
		}
		cur = cur.next
	}
}

func (l *LinkedList[T]) MoveToFront(node *Node[T]) {
	if node == nil || l.head == nil || l.head == node {
		return
	}
	l.Remove(node)
	l.Prepend(node)
}

func (l *LinkedList[T]) MoveToBack(node *Node[T]) {
	if node == nil || l.head == nil || (l.double && l.tail == node) {
		return
	}
	l.Remove(node)
	l.Append(node)
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

	if l.double && l.tail == target {
		return l.Append(node)
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
		if cur == target {
			cur.next = node
			node.next = target
			if l.double {
				target.prev = node
				node.prev = cur
			}
			l.len += 1
			return node
		}
	}

	return nil
}

func (l *LinkedList[T]) PrintNodeForward() {
	fmt.Print("=== PrintNodeForward ===\n")
	cur := l.head
	for cur != nil {
		cur.Print()
		cur = cur.next
		if cur == l.head {
			break
		}
	}
}

func (l *LinkedList[T]) PrintNodeReverse() {
	fmt.Print("=== PrintNodeReverse ===\n")
	if !l.double {
		fmt.Print("list is not doubly linked\n\n")
		return
	}
	cur := l.tail
	for cur != nil {
		cur.Print()
		cur = cur.prev
		if cur == l.tail {
			break
		}
	}
}

// func InsertBeforeValue[T any, E comparable](l *LinkedList[T], value T, target E, fn ExtractFn[T, E]) *Node[T] {
// 	node := &Node[T]{
// 		Value: value,
// 		list:  l,
// 	}

// 	if fn(l.head) == target {
// 		return l.Prepend(value)
// 	}

// 	cur := l.head
// 	for cur.next != nil {
// 		if fn(cur.next) == target {
// 			move := cur.next
// 			cur.next = node
// 			node.next = move
// 			if l.double && move != nil {
// 				move.prev = node
// 				node.prev = cur
// 			}
// 			l.len += 1
// 			return node
// 		}
// 		cur = cur.next
// 	}

// 	return nil
// }

// func InsertAfterValue[T any, E comparable](l *LinkedList[T], value T, target E, fn ExtractFn[T, E]) *Node[T] {
// 	node := &Node[T]{
// 		Value: value,
// 		list:  l,
// 	}

// 	if fn(l.head) == target {
// 		move := l.head.next
// 		l.head.next = node
// 		node.next = move
// 		if l.double && move != nil {
// 			move.prev = node
// 		}
// 		l.len += 1
// 		return node
// 	}

// 	if l.double && fn(l.tail) == target {
// 		return l.Append(value)
// 	}

// 	cur := l.head
// 	for cur.next != nil {
// 		cur = cur.next
// 		if fn(cur) == target {
// 			move := cur.next
// 			cur.next = node
// 			node.next = move
// 			if l.double && move != nil {
// 				move.prev = node
// 				node.prev = cur
// 			}
// 			l.len += 1
// 			return node
// 		}
// 	}

// 	return nil
// }

// func MoveToFrontByValue[T any, E comparable](l *LinkedList[T], target E, fn ExtractFn[T, E]) {
// 	if l.head == nil {
// 		return
// 	}

// 	node := FindNodeByValue(l, target, fn)
// 	if node == nil {
// 		return
// 	}

// 	if node == l.head {
// 		return
// 	}

// 	old := l.head
// 	l.head = node
// 	node.next = old

// 	if l.double {
// 		old.prev = l.head
// 		if l.circular {
// 			l.head.prev = l.tail
// 		}
// 	}
// }

// func MoveToBackByValue[T any, E comparable](l *LinkedList[T], target E, fn ExtractFn[T, E]) {
// 	if l.head == nil {
// 		return
// 	}

// 	node := FindNodeByValue(l, target, fn)
// 	if node == nil {
// 		return
// 	}

// 	if l.double {
// 		if l.tail == node {
// 			return
// 		}

// 		old := l.tail
// 		l.tail = node
// 		node.prev = old
// 		old.next = l.tail
// 		if l.circular {
// 			l.tail.next = l.head
// 		}
// 		return
// 	}

// 	cur := l.head
// 	for cur.next != nil {
// 		if cur.next == l.head {
// 			break
// 		}
// 		cur = cur.next
// 	}
// 	cur.next = node
// 	if l.circular {
// 		node.next = l.head
// 	} else {
// 		node.next = nil
// 	}
// }

// func MoveBeforeValue[T any, E comparable](l *LinkedList[T], node *Node[T], target E, fn ExtractFn[T, E]) *LinkedList[T] {

// 	return l
// }

// func MoveAfterValue[T any, E comparable](l *LinkedList[T], node *Node[T], target E, fn ExtractFn[T, E]) *LinkedList[T] {

// 	return l
// }

// func (l *LinkedList[T]) RemoveNode(node *Node[T]) {
// 	if l.head == nil {
// 		return
// 	}

// 	switch node {
// 	case nil:
// 		return
// 	case l.head:
// 		println("HEAD")
// 		l.head = node.next
// 		l.len -= 1

// 		if l.len == 0 {
// 			l.head = nil
// 			l.tail = nil
// 			return
// 		}

// 		if !l.double && l.circular {

// 		}

// 		if l.double && l.circular {
// 			l.head.prev = l.tail
// 			l.tail.next = l.head
// 		}

// 	case l.tail:
// 		println("TAIL")
// 		l.tail = node.prev
// 		l.len -= 1

// 		if l.len == 0 {
// 			l.head = nil
// 			l.tail = nil
// 			return
// 		}

// 		if l.circular {
// 			l.head.prev = l.tail
// 			l.tail.next = l.head
// 		}

// 	default:
// 		println("MIDDLE")
// 		cur := l.head
// 		for cur.next != nil {
// 			if cur.next == node {
// 				cur.next = node.next
// 				l.len -= 1

// 				if l.double {
// 					cur.next.prev = cur
// 				}

// 				return
// 			}
// 			cur = cur.next
// 		}
// 	}
// }

// func ArrToSLL[T any](arr []T) *SLL[T] {
// 	var head, tail *Node[T]

// 	for _, v := range arr {
// 		node := &Node[T]{
// 			Data: v,
// 			next:  nil,
// 		}

// 		if head == nil {
// 			head = node
// 			tail = node
// 		} else {
// 			tail.next = node
// 			tail = node
// 		}
// 	}

// 	return &SLL[T]{head: head}
// }
