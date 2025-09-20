package list

import "fmt"

type NodeS[T any] struct {
	Value T
	next  *NodeS[T]
	list  *SinglyLinkedList[T]
}

func (n *NodeS[T]) PrintNode() {
	var next *T
	if n.next != nil {
		next = &n.next.Value
	}

	place := "NODE"
	if n == n.list.head {
		place = "HEAD"
	}
	if n == n.list.tail {
		place = "TAIL"
	}

	fmt.Printf("%s: Value %v next %v\n", place, n.Value, next)
}

type extractFnS[T any, E comparable] func(node *NodeS[T]) E

func FindNodeSByValue[T any, E comparable](l *SinglyLinkedList[T], target E, fn extractFnS[T, E]) *NodeS[T] {
	if l.head == nil {
		return nil
	}

	if fn(l.head) == target {
		return l.head
	}

	if fn(l.tail) == target {
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

func (n *NodeS[T]) Next() *NodeS[T] {
	if n.next != nil && n.list != nil && n.next != n.list.head {
		return n.next
	}
	return nil
}
