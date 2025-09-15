package list

import "fmt"

type Node[T any] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
	list  *LinkedList[T]
}

func (n *Node[T]) Next() *Node[T] {
	if n.next != nil && n.list != nil && n.next != n.list.head {
		return n.next
	}
	return nil
}

func (n *Node[T]) Prev() *Node[T] {
	if !n.list.double {
		return nil
	}
	if n.prev != nil && n.list != nil && n.prev != n.list.tail {
		return n.prev
	}
	return nil
}

func (n *Node[T]) Print() {
	var next, prev T
	if n.next != nil {
		next = n.next.Value
	}
	if n.prev != nil {
		prev = n.prev.Value
	}
	fmt.Printf("NODE: Value %v next %v prev %v\n", n.Value, next, prev)
}

type extractFn[T any, E comparable] func(node *Node[T]) E

func FindNodeByValue[T any, E comparable](l *LinkedList[T], target E, fn extractFn[T, E]) *Node[T] {
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
