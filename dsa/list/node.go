package list

import "fmt"

type Node[T any] struct {
	Data T
	next *Node[T]
	prev *Node[T]
	list *LinkedList[T]
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
		next = n.next.Data
	}
	if n.prev != nil {
		prev = n.prev.Data
	}
	fmt.Printf("NODE: Data %v next %v prev %v\n\n", n.Data, next, prev)
}
