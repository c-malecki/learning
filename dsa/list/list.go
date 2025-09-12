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
	fmt.Printf("LIST: head %+v tail %+v len %v double %v circular %v\n\n", l.head, l.tail, l.len, l.double, l.circular)
}

// insert (after,before)
// move (after,before,front,back), sort

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

func (l *LinkedList[T]) InsertBack(data T) *Node[T] {
	node := &Node[T]{
		Data: data,
		list: l,
	}

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

func (l *LinkedList[T]) InsertFront(data T) *Node[T] {
	node := &Node[T]{
		Data: data,
		next: l.head,
		list: l,
	}

	if l.circular {
		return l.insertFrontCircular(node)
	}
	return l.insertFrontLinear(node)
}

func (l *LinkedList[T]) insertFrontLinear(node *Node[T]) *Node[T] {
	if l.head == nil {
		l.head = node
		l.len += 1
		return node
	}

	old := l.head
	l.head = node
	node.next = old

	if l.double {
		old.prev = l.head
	}

	l.len += 1

	return node
}

func (l *LinkedList[T]) insertFrontCircular(node *Node[T]) *Node[T] {
	if l.head == nil {
		l.head = node
		l.head.next = l.head
		l.head.prev = l.head
		l.len += 1
		return node
	}

	old := l.head
	node.next = old
	l.head = node

	if l.double {
		old.prev = l.head
		l.head.prev = l.tail
		l.tail.next = l.head
	}

	l.len += 1

	return node
}

type ExtractFn[T any, E comparable] func(node *Node[T]) E

func Remove[T any, E comparable](l *LinkedList[T], data E, fn ExtractFn[T, E]) {
	if l.double {
		removeDoubly(l, data, fn)
	} else {
		removeSingly(l, data, fn)
	}
}

func removeSingly[T any, E comparable](l *LinkedList[T], data E, fn ExtractFn[T, E]) {
	if l.head == nil {
		return
	}

	if fn(l.head) == data {
		old := l.head
		l.head = l.head.next
		if l.circular {
			cur := l.head
			for cur.next != nil {
				if cur.next == old {
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
		if fn(cur.next) == data {
			cur.next = cur.next.next
			l.len -= 1
			break
		}
		cur = cur.next
	}
}

func removeDoubly[T any, E comparable](l *LinkedList[T], data E, fn ExtractFn[T, E]) {
	if l.head == nil {
		return
	}

	if fn(l.head) == data {
		l.head = l.head.next
		if l.circular {
			l.head.prev = l.tail
			l.tail.next = l.head
		} else {
			l.head.prev = nil
		}
		l.len -= 1
		return
	}

	cur := l.head
	for cur.next != nil {
		if fn(cur.next) == data {
			cur.next = cur.next.next
			cur.next.prev = cur
			l.len -= 1
			break
		}
		cur = cur.next
	}
}

func InsertBefore[T any, E comparable](l *LinkedList[T], data T, target E, fn ExtractFn[T, E]) *Node[T] {
	node := &Node[T]{
		Data: data,
		list: l,
	}

	if fn(l.head) == target {
		return l.InsertFront(data)
	}

	cur := l.head
	for cur.next != nil {
		if fn(cur.next) == target {
			move := cur.next
			cur.next = node
			node.next = move
			if l.double && move != nil {
				move.prev = node
				node.prev = cur
			}
			l.len += 1
			return node
		}
		cur = cur.next
	}

	return nil
}

func InsertAfter[T any, E comparable](l *LinkedList[T], data T, target E, fn ExtractFn[T, E]) *Node[T] {
	node := &Node[T]{
		Data: data,
		list: l,
	}

	if fn(l.head) == target {
		move := l.head.next
		l.head.next = node
		node.next = move
		if l.double && move != nil {
			move.prev = node
		}
		l.len += 1
		return node
	}

	if l.double && fn(l.tail) == target {
		return l.InsertBack(data)
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
		if fn(cur) == target {
			move := cur.next
			cur.next = node
			node.next = move
			if l.double && move != nil {
				move.prev = node
				node.prev = cur
			}
			l.len += 1
			return node
		}
	}

	return nil
}

func MoveBefore[T any, E comparable](l *LinkedList[T], node *Node[T], target E, fn ExtractFn[T, E]) *LinkedList[T] {

	return l
}

func MoveAfter[T any, E comparable](l *LinkedList[T], node *Node[T], target E, fn ExtractFn[T, E]) *LinkedList[T] {

	return l
}

func FindNode[T any, E comparable](l *LinkedList[T], target E, fn ExtractFn[T, E]) *Node[T] {
	if l.head == nil {
		return nil
	}

	if fn(l.head) == target {
		return l.head
	}

	if l.double && fn(l.tail) == target {
		return l.tail
	}

	cur := l.head
	for cur.next != nil && cur.next != l.head {
		cur = cur.next
		if fn(cur) == target {
			return cur
		}
	}

	return nil
}

func MoveFront[T any, E comparable](l *LinkedList[T], target E, fn ExtractFn[T, E]) {
	if l.head == nil {
		return
	}

	node := FindNode(l, target, fn)
	if node == nil {
		return
	}

	if node == l.head {
		return
	}

	old := l.head
	l.head = node
	node.next = old

	if l.double {
		old.prev = l.head
		if l.circular {
			l.head.prev = l.tail
		}
	}
}

func MoveBack[T any, E comparable](l *LinkedList[T], target E, fn ExtractFn[T, E]) {
	if l.head == nil {
		return
	}

	node := FindNode(l, target, fn)
	if node == nil {
		return
	}

	if l.double {
		if l.tail == node {
			return
		}

		old := l.tail
		l.tail = node
		node.prev = old
		old.next = l.tail
		if l.circular {
			l.tail.next = l.head
		}
		return
	}

	cur := l.head
	for cur.next != nil {
		if cur.next == l.head {
			break
		}
		cur = cur.next
	}
	cur.next = node
	if l.circular {
		node.next = l.head
	} else {
		node.next = nil
	}
}

func (l *LinkedList[T]) PrintNodeForward() {
	fmt.Print("PRINT NODES FORWARD\n")
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
	fmt.Print("PRINT NODES REVERSE\n")
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
